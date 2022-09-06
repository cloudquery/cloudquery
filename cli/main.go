package main

import (
	"os"
	"time"

	"github.com/cloudquery/cloudquery/cli/cmd"
	"github.com/getsentry/sentry-go"
)

const sentryFlushTimeout = 5 * time.Second

func main() {
	if err := cmd.NewCmdRoot().Execute(); err != nil {
		sentry.CaptureMessage(err.Error())
		sentry.Flush(sentryFlushTimeout)
		os.Exit(1)
	}
	sentry.Flush(sentryFlushTimeout)
}
