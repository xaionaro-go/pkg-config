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
	case "--cflags":
		runPkgConfig(os.Args[1:]...)
	case "--libs":
		var out bytes.Buffer
		cmd := exec.Command("pkg-config", append([]string{"pkg-config", "--static"}, os.Args[1:]...)...)
		cmd.Stdout = &out
		cmd.Run()
		if cmd.ProcessState.ExitCode() != 0 {
			fmt.Fprintf(os.Stdout, "%s", out.String())
			os.Exit(cmd.ProcessState.ExitCode())
		}
		var result []string
		for _, w := range strings.Split(strings.Trim(out.String(), "\n"), " ") {
			switch {
			case strings.HasPrefix(w, "-f"):
				continue
			}
			result = append(result, w)
		}
		fmt.Fprintf(os.Stdout, "-Wl,-Bstatic %s -Wl,-Bdynamic\n", strings.Join(result, " "))
		os.Exit(0)
	default:
		panic(fmt.Errorf("%v", os.Args))
	}
}

func runPkgConfig(args ...string) {
	cmd := exec.Command("pkg-config", append([]string{"pkg-config"}, args...)...)
	cmd.Run()
	os.Exit(cmd.ProcessState.ExitCode())
}
