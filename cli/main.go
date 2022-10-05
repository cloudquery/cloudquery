package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/cloudquery/cloudquery/cli/cmd"
	"github.com/getsentry/sentry-go"
)

const sentryFlushTimeout = 5 * time.Second

func main() {
	defer func() {
		err := recover()
		if err != nil {
			originalMessage := fmt.Sprintf("panic: %v\n\n%s", err, string(debug.Stack()))
			sentry.CurrentHub().CaptureMessage(originalMessage)
			panic(err)
		}
	}()

	if err := cmd.NewCmdRoot().Execute(); err != nil {
		// This is fine that the defer function is not being calles as it means there was no panic
		//nolint:gocritic
		os.Exit(1)
	}
	sentry.Flush(sentryFlushTimeout)
}
