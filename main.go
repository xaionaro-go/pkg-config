package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/facebookincubator/go-belt"
	"github.com/facebookincubator/go-belt/tool/logger"
	xlogrus "github.com/facebookincubator/go-belt/tool/logger/implementation/logrus"
	"github.com/sirupsen/logrus"
	"github.com/xaionaro-go/pkg-config-wrapper/pkg/consts"
	"github.com/xaionaro-go/pkg-config-wrapper/pkg/pkgconfig"
)

func ctx() context.Context {
	ctx := context.Background()
	ll := xlogrus.DefaultLogrusLogger()
	l := xlogrus.New(ll).WithLevel(logger.LevelTrace)
	ctx = logger.CtxWithLogger(ctx, l)

	if !func() bool {
		logFilePath := os.Getenv(consts.EnvVarLogFile)
		if logFilePath == "" {
			return false
		}

		logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
		if err != nil {
			l.Errorf("unable to open log-file '%s': %v", logFilePath, err)
			return false
		}

		ll.SetOutput(logFile)
		return true
	}() {
		ll.Formatter.(*logrus.TextFormatter).ForceColors = true
	}
	return ctx
}

func main() {
	defer os.Stderr.Sync()
	defer os.Stdout.Sync()
	if len(os.Args) < 2 {
		panic("not enough arguments")
	}

	ctx := ctx()
	defer belt.Flush(ctx)

	var opts pkgconfig.Options

	erasePatterns := parseList(os.Getenv(consts.EnvVarEraseList))
	if len(erasePatterns) > 0 {
		opts = append(opts, pkgconfig.OptionErasePatterns(erasePatterns))
	}

	staticLibsPatterns := parseList(os.Getenv(consts.EnvVarStaticLibsList))
	if len(staticLibsPatterns) > 0 {
		opts = append(opts, pkgconfig.OptionForceStaticLinkPatterns(staticLibsPatterns))
	}

	dynamicLibsPatterns := parseList(os.Getenv(consts.EnvVarDynamicLibsList))
	if len(dynamicLibsPatterns) > 0 {
		opts = append(opts, pkgconfig.OptionForceDynamicLinkPatterns(dynamicLibsPatterns))
	}

	pkgConfig := pkgconfig.NewPkgConfig(opts...)
	result, errorMsg, exitCode, err := pkgConfig.Run(ctx, os.Args[1:]...)
	if _, err := fmt.Fprintf(os.Stderr, "%s", errorMsg); err != nil {
		panic(err)
	}
	if _, err := fmt.Fprintf(os.Stdout, "%s\n", strings.Join(result, " ")); err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	os.Exit(exitCode)
}

func parseList(
	input string,
) pkgconfig.Patterns {
	var result pkgconfig.Patterns
	for _, w := range strings.Split(input, ",") {
		word := strings.Trim(w, " ")
		if len(word) == 0 {
			continue
		}
		result = append(result, pkgconfig.Pattern(word))
	}
	return result
}
