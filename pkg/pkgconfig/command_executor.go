package pkgconfig

import (
	"bytes"
	"os/exec"
)

type CommandExecutor interface {
	Execute(cmd string, args ...string) ([]byte, []byte, int, error)
}

var DefaultCommandExecutor = &RealCommandExecutor{}

type RealCommandExecutor struct{}

func (RealCommandExecutor) Execute(
	arg0 string,
	args ...string,
) ([]byte, []byte, int, error) {
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd := exec.Command(pkgConfig, args...)
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		return stdOut.Bytes(), stdErr.Bytes(), -1, err
	}
	return stdOut.Bytes(), stdErr.Bytes(), cmd.ProcessState.ExitCode(), err
}
