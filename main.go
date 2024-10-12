package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/xaionaro-go/pkg-config/pkg/consts"
	"github.com/xaionaro-go/pkg-config/pkg/pkgconfig"
)

func main() {
	if len(os.Args) < 2 {
		panic("not enough arguments")
	}

	var opts pkgconfig.Options

	staticLibsPatterns := parseList(os.Getenv(consts.EnvVarStaticLibsList))
	if len(staticLibsPatterns) > 0 {
		opts = append(opts, pkgconfig.OptionForceStaticLinkPatterns(staticLibsPatterns))
	}

	dynamicLibsPatterns := parseList(os.Getenv(consts.EnvVarDynamicLibsList))
	if len(dynamicLibsPatterns) > 0 {
		opts = append(opts, pkgconfig.OptionForceDynamicLinkPatterns(dynamicLibsPatterns))
	}

	pkgConfig := pkgconfig.NewPkgConfig(opts...)
	result, errorMsg, exitCode, err := pkgConfig.Run(os.Args[1:]...)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintf(os.Stderr, "%s", errorMsg)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintf(os.Stdout, "%s\n", strings.Join(result, " "))
	if err != nil {
		panic(err)
	}
	os.Exit(exitCode)
}

func parseList(
	input string,
) []string {
	var result []string
	for _, w := range strings.Split(input, ",") {
		word := strings.Trim(w, " ")
		if len(word) == 0 {
			continue
		}
		result = append(result, word)
	}
	return result
}
