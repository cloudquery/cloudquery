package cmd

import (
	"context"
	"os"
	"time"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/internal/signalcontext"
	"github.com/cloudquery/cloudquery/internal/telemetry"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cloudquery/pkg/ui/console"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func handleCommand(f func(context.Context, *console.Client, *cobra.Command, []string) error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		tele := telemetry.New(cmd.Context(), telemetryOpts()...)

		tracer := tele.Tracer()
		ctx, span := tracer.Start(cmd.Context(),
			"cli:"+cmd.CommandPath(),
			trace.WithAttributes(
				attribute.String("command", cmd.CommandPath()),
			),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		ender := func() {
			span.End(
				trace.WithStackTrace(false),
			)
			tele.Shutdown(cmd.Context())
		}
		defer ender()

		if err := handleConsole(ctx, tele, cmd, args, f); err != nil {
			if ee, ok := err.(*console.ExitCodeError); ok {
				ender() // err is not recorded
				os.Exit(ee.ExitCode)
			}

			tele.RecordError(span, err)
			ender()

			cmd.PrintErrln(err)
			os.Exit(1)
		}
	}
}

func handleConsole(ctx context.Context, tele *telemetry.Client, cmd *cobra.Command, args []string, f func(context.Context, *console.Client, *cobra.Command, []string) error) error {
	configPath := viper.GetString("configPath")

	ctx, _ = signalcontext.WithInterrupt(ctx, logging.NewZHcLog(&log.Logger, ""))
	var c *console.Client

	delayMessage := ui.IsTerminal()

	switch cmd.Name() {
	// Don't init console client with these commands
	case "completion", "options":
		delayMessage = false
	default:
		var err error
		c, err = console.CreateClient(ctx, configPath)
		if err != nil {
			return err
		}
		defer c.Client().Close()
	}

	if tele.NewCookie() {
		ui.ColorizedOutput(ui.ColorInfo, "Anonymous telemetry collection enabled. Run with --no-telemetry to disable, or check docs at https://docs.cloudquery.io/docs/telemetry\n")
		if delayMessage {
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
