package pkgconfig

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockCommandExecutor struct {
	ExecuteFunc func(
		cmd string,
		args ...string,
	) ([]byte, []byte, int, error)
}

var _ CommandExecutor = (*mockCommandExecutor)(nil)

func (e *mockCommandExecutor) Execute(
	cmd string,
	args ...string,
) ([]byte, []byte, int, error) {
	return e.ExecuteFunc(cmd, args...)
}

func TestPkgConfigRun(t *testing.T) {
	callCount := 0
	pkgConfig := NewPkgConfig(
		OptionCommandExecutor{&mockCommandExecutor{
			ExecuteFunc: func(cmd string, args ...string) ([]byte, []byte, int, error) {
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
					require.Equal(t, []string{"--random-arg", "--libs-only-l", "libm"}, args, callCount)
					return []byte("-lm"), nil, 0, nil
				default:
					return nil, nil, -1, fmt.Errorf("the command executor was called too many times")
				}
			},
		}},
		OptionForceStaticLinkPatterns([]string{"libav*"}),
		OptionForceDynamicLinkPatterns([]string{"libvlc", "libandroid"}),
	)
	output, errMsg, exitCode, err := pkgConfig.Run(
		"--random-arg", "--libs-only-l", "libpthread", "libm", "libavcodec", "libvlc",
	)
	require.NoError(t, err)
	require.Equal(t, 0, exitCode)
	require.Empty(t, errMsg, fmt.Sprintf("%X", errMsg))
	require.Equal(t, []string{
		"-lm", "-Wl,-Bstatic", "-lpthread", "-lx264", "-lavcodec", "-Wl,-Bdynamic", "-lvlc", "-landroid",
	}, output)
	require.Equal(t, 3, callCount)
}
