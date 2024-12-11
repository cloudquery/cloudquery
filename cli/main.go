package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"

	"github.com/cloudquery/cloudquery/cli/v6/cmd"
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func executeRootCmdWithContext() error {
	// trap Ctrl+C and other signals, then call cancel on the context
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer func() {
		signal.Stop(c)
	}()
	var gotSignal os.Signal
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case gotSignal = <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	err := cmd.NewCmdRoot().ExecuteContext(ctx)
	cancel()
	wg.Wait()
	if gotSignal != nil && err != nil {
		err = fmt.Errorf("received %v signal from OS: %w", gotSignal.String(), err)
	} else if gotSignal != nil {
		err = fmt.Errorf("received %v signal from OS", gotSignal.String())
	}
	return err
}

func main() {
	exitCode := 0
	defer func() {
		err := recover()
		if err != nil {
			originalMessage := fmt.Sprintf("panic: %v\n\n%s", err, string(debug.Stack()))
			sentry.CurrentHub().CaptureMessage(originalMessage)
			panic(err)
		}

		os.Exit(exitCode)
	}()
	defer cmd.CloseLogFile()

	// This ensures we don't print anything until logging is configured
	log.Logger = log.Level(zerolog.Disabled)
	if err := executeRootCmdWithContext(); err != nil {
		log.Error().Err(err).Msg("exiting with error")
		exitCode = 1
	}
}
