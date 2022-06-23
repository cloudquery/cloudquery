package cmd

import (
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/errors"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/spf13/cobra"
)

var (
	providerHelpMsg = "Top-level command to interact with providers."
	providerCmd     = &cobra.Command{
		Use:   "provider [subcommand]",
		Short: providerHelpMsg,
		Long:  providerHelpMsg,
		Example: `
  # Downloads all providers specified in config.hcl:
  cloudquery provider download
  # Sync (Upgrade or Downgrade) all providers specified in config.hcl This will also create the schema.
  cloudquery provider sync 
  # Sync one or more providers
  cloudquery provider sync aws, gcp
  # Drop provider schema, running fetch again will recreate all tables unless --skip-build-tables is specified
  cloudquery provider drop aws
`,
		Version: core.Version,
	}

	providerSyncHelpMsg = "Download the providers specified in config.hcl and re-create their database schema"
	providerSyncCmd     = &cobra.Command{
		Use:   "sync [providers,...]",
		Short: providerSyncHelpMsg,
		Long:  providerSyncHelpMsg,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), getConfigFile(), false, nil, instanceId)
			if err != nil {
				return err
			}
			_, diags := c.SyncProviders(cmd.Context(), args...)
			errors.CaptureDiagnostics(diags, map[string]string{"command": "provider_sync"})
			if diags.HasErrors() {
				return fmt.Errorf("failed to sync providers")
			}
			return nil
		},
	}

	providerForce       bool
	providerDropHelpMsg = "Drops provider schema from database"
	providerDropCmd     = &cobra.Command{
		Use:   "drop [provider]",
		Short: providerDropHelpMsg,
		Long:  providerDropHelpMsg,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), getConfigFile(), false, nil, instanceId)
			if err != nil {
				return err
			}
			if !providerForce {
				ui.ColorizedOutput(ui.ColorWarning, "WARNING! This will drop all tables for the given provider. If you wish to continue, use the --force flag.\n")
				return diag.FromError(fmt.Errorf("if you wish to continue, use the --force flag"), diag.USER)
			}
			diags := c.DropProvider(cmd.Context(), args[0])
			errors.CaptureDiagnostics(diags, map[string]string{"command": "provider_drop"})
			if diags.HasErrors() {
				return fmt.Errorf("failed to drop provider %s", args[0])
			}
			return nil
		},
	}

	providerDownloadHelpMsg = "Downloads all providers specified in config.hcl."
	providerDownloadCmd     = &cobra.Command{
		Use:   "download",
		Short: providerDownloadHelpMsg,
		Long:  providerDownloadHelpMsg,
		Example: `
  # Downloads all providers specified in config.hcl:
  ./cloudquery provider download
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), getConfigFile(), false, nil, instanceId)
			if err != nil {
				return err
			}
			diags := c.DownloadProviders(cmd.Context())
			errors.CaptureDiagnostics(diags, map[string]string{"command": "provider_download"})
			if diags.HasErrors() {
				return fmt.Errorf("failed to download providers")
			}
			return nil
		},
	}

	lastUpdate                 time.Duration
	dryRun                     bool
	providerRemoveStaleHelpMsg = "Remove stale resources from one or more providers in database"
	providerRemoveStaleCmd     = &cobra.Command{
		Use:   "purge [provider]",
		Short: providerRemoveStaleHelpMsg,
		Long:  providerRemoveStaleHelpMsg,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), getConfigFile(), false, nil, instanceId)
			if err != nil {
				return err
			}
			diags := c.RemoveStaleData(cmd.Context(), lastUpdate, dryRun, args)
			errors.CaptureDiagnostics(diags, map[string]string{"command": "provider_purge"})
			if diags.HasErrors() {
				return fmt.Errorf("failed to remove stale data")
			}
			return nil
		},
	}
)

func init() {
	providerRemoveStaleCmd.Flags().DurationVar(&lastUpdate, "last-update", time.Hour*1,
		"last-update is the duration from current time we want to remove resources from the database. "+
			"For example 24h will remove all resources that were not update in last 24 hours. Duration is a string with optional unit suffix such as \"2h45m\" or \"7d\"")
	providerRemoveStaleCmd.Flags().BoolVar(&dryRun, "dry-run", true, "")
	providerDropCmd.Flags().BoolVar(&providerForce, "force", false, "Really drop tables for the provider")
	providerCmd.AddCommand(providerDownloadCmd, providerSyncCmd, providerDropCmd, providerRemoveStaleCmd)
	providerCmd.SetUsageTemplate(usageTemplateWithFlags)
	rootCmd.AddCommand(providerCmd)
}
