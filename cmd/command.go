package cmd

import (
	"context"
	"os"
	"time"

	"github.com/cloudquery/cloudquery/internal/analytics"
	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/internal/signalcontext"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// fileDescriptorF gets set trough system relevant files like ulimit_unix.go
var fileDescriptorF func()

func handleCommand(f func(context.Context, *console.Client, *cobra.Command, []string) error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		var exitWithCode int
		defer func() {
			if exitWithCode > 0 {
				os.Exit(exitWithCode)
			}
		}()
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			sentry.CurrentHub().Recover(err)
			panic(err)
		}()

		if err := handleConsole(cmd.Context(), cmd, args, f); err != nil {
			if dd := diag.FromError(err, diag.INTERNAL); dd.HasErrors() { // Check if already diags
				exitWithCode = 1
			}
			cmd.PrintErrln(err)
		}
	}
}

func handleConsole(ctx context.Context, cmd *cobra.Command, args []string, f func(context.Context, *console.Client, *cobra.Command, []string) error) error {

	cfgPath := viper.GetString("configPath")
	ctx, _ = signalcontext.WithInterrupt(ctx, logging.NewZHcLog(&log.Logger, ""))
	var (
		c            *console.Client
		cfgMutator   func(*config.Config) error
		telemetryMsg = true
	)

	switch cmd.Name() {
	// Don't init console client with these commands
	case "completion", "options":
		telemetryMsg = false
	case "init":
		// No console client created here
	case "describe":
		var err error
		c, err = console.CreateClient(ctx, cfgPath, true, cfgMutator, instanceId)
		if err != nil {
			return err
		}
	case "fetch":
		cfgMutator = filterConfigProviders(args)
		fallthrough
	default:
		var err error
		c, err = console.CreateClient(ctx, cfgPath, false, cfgMutator, instanceId)
		if err != nil {
			return err
		}
	}

	if telemetryMsg && analytics.Enabled() {
		ui.ColorizedOutput(ui.ColorInfo, "Anonymous telemetry collection and crash reporting enabled. Run with --no-telemetry to disable, or check docs at https://docs.cloudquery.io/docs/cli/telemetry\n")
		if ui.IsTerminal() {
			select {
			case <-time.After(2 * time.Second):
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	if err := f(ctx, c, cmd, args); err != nil {
		return err
	}

	return nil
}
