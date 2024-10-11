package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic("not enough arguments")
	}
	switch os.Args[1] {
	case "--libs":
		var out bytes.Buffer
		cmd := exec.Command("pkg-config", append([]string{"pkg-config"}, os.Args[1:]...)...)
		cmd.Stdout = &out
		cmd.Run()
		if cmd.ProcessState.ExitCode() != 0 {
			fmt.Fprintf(os.Stdout, "%s", out.String())
			os.Exit(cmd.ProcessState.ExitCode())
		}
		fmt.Fprintf(os.Stdout, "-Wl,-Bstatic %s -Wl,-Bdynamic\n", strings.Trim(out.String(), "\n"))
		os.Exit(0)
	default:
		cmd := exec.Command("pkg-config", append([]string{"pkg-config"}, os.Args[1:]...)...)
		cmd.Run()
		os.Exit(cmd.ProcessState.ExitCode())
	}
}
