package pkgconfig

import (
	"context"
	"fmt"
	"testing"

	"github.com/facebookincubator/go-belt/tool/logger"
	xlogrus "github.com/facebookincubator/go-belt/tool/logger/implementation/logrus"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

type mockCommandExecutor struct {
	ExecuteFunc func(
		ctx context.Context,
		cmd string,
		args ...string,
	) ([]byte, []byte, int, error)
}

var _ CommandExecutor = (*mockCommandExecutor)(nil)

func (e *mockCommandExecutor) Execute(
	ctx context.Context,
	cmd string,
	args ...string,
) ([]byte, []byte, int, error) {
	return e.ExecuteFunc(ctx, cmd, args...)
}

func ctx() context.Context {
	ctx := context.Background()
	ll := xlogrus.DefaultLogrusLogger()
	ll.Formatter.(*logrus.TextFormatter).ForceColors = true
	l := xlogrus.New(ll).WithLevel(logger.LevelTrace)
	return logger.CtxWithLogger(ctx, l)
}

func TestPkgConfigRun(t *testing.T) {
	ctx := ctx()
	callCount := 0
	pkgConfig := NewPkgConfig(
		OptionCommandExecutor{&mockCommandExecutor{
			ExecuteFunc: func(_ context.Context, cmd string, args ...string) ([]byte, []byte, int, error) {
				if cmd != pkgConfig {
					return nil, nil, -1, fmt.Errorf("unexpected command '%s' (expected: '%s')", cmd, pkgConfig)
				}

				callCount++
				switch callCount {
				case 1:
					require.Equal(t, []string{"--static", "--random-arg", "--libs-only-l", "libavcodec"}, args, callCount)
					return []byte("-lpthread -lx264 -lavcodec -landroid"), []byte{}, 0, nil
				case 2:
					require.Equal(t, []string{"--shared", "--random-arg", "--libs-only-l", "libvlc", "libandroid"}, args, callCount)
					return []byte("-lvlc -landroid"), []byte{}, 0, nil
				case 3:
					require.Equal(t, []string{"--random-arg", "--libs-only-l", "libm", "librandom"}, args, callCount)
					return []byte("-lm -lrandom"), nil, 0, nil
				default:
					return nil, nil, -1, fmt.Errorf("the command executor was called too many times")
				}
			},
		}},
		OptionErasePatterns{"-lrandom"},
		OptionForceStaticLinkPatterns{"libav*"},
		OptionForceDynamicLinkPatterns{"libvlc", "libandroid"},
	)
	output, errMsg, exitCode, err := pkgConfig.Run(
		ctx,
		"--random-arg", "--libs-only-l", "libpthread", "libm", "librandom", "libavcodec", "libvlc",
	)
	require.NoError(t, err)
	require.Equal(t, 0, exitCode)
	require.Empty(t, errMsg, fmt.Sprintf("%X", errMsg))
	require.Equal(t, []string{
		"-lm", "-Wl,-Bstatic", "-lpthread", "-lx264", "-lavcodec", "-Wl,-Bdynamic", "-lvlc", "-landroid",
	}, output)
	require.Equal(t, 3, callCount)
}
