package provider

import (
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/cmd/utils"
	"github.com/cloudquery/cloudquery/pkg/errors"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
)

var (
	lastUpdate time.Duration
	dryRun     bool
	purgeShort = "Remove stale resources from one or more providers in database"
)

func newCmdProviderPurge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "purge [provider]",
		Short: purgeShort,
		Long:  purgeShort,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), utils.GetConfigFile(), false, nil, utils.InstanceId)
			if err != nil {
				return err
			}
			diags := c.RemoveStaleData(cmd.Context(), lastUpdate, dryRun, args)
			errors.CaptureDiagnostics(diags, map[string]string{"command": "provider_purge"})
			if diags.HasErrors() {
				return fmt.Errorf("failed to remove stale data: %w", diags)
			}
			return nil
		},
	}
	cmd.Flags().DurationVar(&lastUpdate, "last-update", time.Hour*1,
		"last-update is the duration from current time we want to remove resources from the database. "+
			"For example 24h will remove all resources that were not update in last 24 hours. Duration is a string with optional unit suffix such as \"2h45m\" or \"7d\"")
	cmd.Flags().BoolVar(&dryRun, "dry-run", true, "")
	return cmd
}
