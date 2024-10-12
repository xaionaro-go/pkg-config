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
					require.Equal(t, []string{"--random-arg", "--libs-only-l", "libandroid"}, args, callCount)
					return []byte("-landroid"), nil, 0, nil
				case 2:
					require.Equal(t, []string{"--static", "--random-arg", "--libs-only-l", "libavcodec"}, args, callCount)
					return []byte("-lm -lavcodec"), []byte{}, 0, nil
				case 3:
					require.Equal(t, []string{"--shared", "--random-arg", "--libs-only-l", "libvlc"}, args, callCount)
					return []byte("-lvlc"), []byte{}, 0, nil
				default:
					return nil, nil, -1, fmt.Errorf("the command executor was called too many times")
				}
			},
		}},
		OptionForceStaticLinkPatterns([]string{"libav*"}),
		OptionForceDynamicLinkPatterns([]string{"libvlc"}),
	)
	output, errMsg, exitCode, err := pkgConfig.Run(
		"--random-arg", "--libs-only-l", "libavcodec", "libvlc", "libandroid",
	)
	require.NoError(t, err)
	require.Equal(t, 0, exitCode)
	require.Empty(t, errMsg, fmt.Sprintf("%X", errMsg))
	require.Equal(t, []string{"-landroid", "-Wl,-Bstatic", "-lm", "-lavcodec", "-Wl,-Bdynamic", "-lvlc"}, output)
	require.Equal(t, 3, callCount)
}
