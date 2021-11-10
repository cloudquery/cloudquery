package cmd

import (
	"os"

	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
)

func handleError(f func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := f(cmd, args); err != nil {
			if ee, ok := err.(*console.ExitCodeError); ok {
				os.Exit(ee.ExitCode)
			}

			cmd.PrintErrln(err)
			os.Exit(1)
		}
	}
}
