package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/cloudquery/cloudquery/pkg/ui/console"
)

func handleError(f func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		tracer := tele.Tracer()
		_, span := tracer.Start(cmd.Context(),
			"cli:"+cmd.CommandPath(),
			trace.WithAttributes(
				attribute.String("command", cmd.CommandPath()),
			),
		)
		ender := func() {
			span.End()
			tele.Shutdown()
		}
		defer ender()

		if err := f(cmd, args); err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			ender()

			if ee, ok := err.(*console.ExitCodeError); ok {
				os.Exit(ee.ExitCode)
			}

			cmd.PrintErrln(err)
			os.Exit(1)
		}
	}
}
