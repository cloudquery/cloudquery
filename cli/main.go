package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"

	"github.com/cloudquery/cloudquery/cli/cmd"
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func executeRootCmdWithContext() error {
	ctx := context.Background()
	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	return cmd.NewCmdRoot().ExecuteContext(ctx)
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			originalMessage := fmt.Sprintf("panic: %v\n\n%s", err, string(debug.Stack()))
			sentry.CurrentHub().CaptureMessage(originalMessage)
			panic(err)
		}
	}()

	// This ensures we don't print anything until logging is configured
	log.Logger = log.Level(zerolog.Disabled)
	if err := executeRootCmdWithContext(); err != nil {
		log.Error().Err(err).Msg("exiting with error")
		//nolint:all This is fine if deferred is not called because there was no panic
		os.Exit(1)
	}
}
