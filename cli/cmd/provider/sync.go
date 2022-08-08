package provider

import (
	"fmt"

	"github.com/cloudquery/cloudquery/cmd/utils"
	"github.com/cloudquery/cloudquery/pkg/errors"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
)

const syncShort = "Download the providers specified in config and re-create their database schema"

func newCmdProviderSync() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync [providers,...]",
		Short: syncShort,
		Long:  syncShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), utils.GetConfigFile(), false, nil, utils.InstanceId)
			if err != nil {
				return err
			}
			_, diags := c.SyncProviders(cmd.Context(), args...)
			errors.CaptureDiagnostics(diags, map[string]string{"command": "provider_sync"})
			if diags.HasErrors() {
				return fmt.Errorf("failed to sync providers %w", diags)
			}
			return nil
		},
	}
	return cmd
}
