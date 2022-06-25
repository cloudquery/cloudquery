package provider

import (
	"fmt"

	"github.com/cloudquery/cloudquery/cmd/util"
	"github.com/cloudquery/cloudquery/pkg/errors"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
)

const (
	syncShort = "Download the providers specified in config and re-create their database schema"
)

func NewCmdProviderSync(parentOptions providerOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync [providers,...]",
		Short: syncShort,
		Long:  syncShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), util.GetConfigFile(parentOptions.Config), false, nil, util.InstanceId)
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
