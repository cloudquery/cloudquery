package provider

import (
	"github.com/spf13/cobra"
)

const syncShort = "Download the providers specified in config and re-create their database schema"

func newCmdProviderSync() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync [providers,...]",
		Short: syncShort,
		Long:  syncShort,
		RunE:  sync,
	}
	return cmd
}

func sync(cmd *cobra.Command, args []string) error {
	return nil
}
