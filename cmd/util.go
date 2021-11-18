package cmd

import (
	"context"
	"os"

	"github.com/cloudquery/cloudquery/internal/telemetry"
	"github.com/cloudquery/cloudquery/pkg/ui/console"

	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func handleError(f func(context.Context, *cobra.Command, []string) error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		tele := telemetry.New(cmd.Context(), telemetryOpts()...)

		tracer := tele.Tracer()
		spanContext, span := tracer.Start(cmd.Context(),
			"cli:"+cmd.CommandPath(),
			trace.WithAttributes(
				attribute.String("command", cmd.CommandPath()),
			),
		)
		ender := func() {
			span.End()
			tele.Shutdown(cmd.Context())
		}
		defer ender()

		if err := f(spanContext, cmd, args); err != nil {
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
