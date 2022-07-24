package provider

import (
	"github.com/spf13/cobra"
)

var (
	providerForce bool
	dropShort     = "Drops provider schema from database"
)

func newCmdProviderDrop() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "drop [provider]",
		Short: dropShort,
		Long:  dropShort,
		Args:  cobra.ExactArgs(1),
		RunE:  drop,
	}
	cmd.Flags().BoolVar(&providerForce, "force", false, "Really drop tables for the provider")
	return cmd
}

func drop(cmd *cobra.Command, args []string) error {
	return nil
}
