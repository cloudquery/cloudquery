package cmd

import (
	"github.com/spf13/cobra"
)

const (
	scaffoldShort = "Create an empty plugin project"
)

func newCmdScaffold() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scaffold",
		Short: scaffoldShort,
	}
	cmd.AddCommand(newCmdScaffoldSource(), newCmdScaffoldDestination())
	return cmd
}
