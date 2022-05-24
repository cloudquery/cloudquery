package cmd

import (
	"context"
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
  # Upgrades all providers specified in config.hcl
  cloudquery provider upgrade 
  # Upgrade one or more providers
  cloudquery provider upgrade aws
  # Downgrades all providers specified in config.hcl
  cloudquery provider downgrade 
  # Downgrades one or more providers
  cloudquery provider downgrade aws, gcp
  # Drop provider schema, running fetch again will recreate all tables unless --skip-build-tables is specified
  cloudquery provider drop aws
  # build provider schema
  cloudquery provider build-schema aws
`,
		Version: core.Version,
	}

	providerUpgradeHelpMsg = "Upgrades one or more providers schema version based on config.hcl"
	providerUpgradeCmd     = &cobra.Command{
		Use:   "upgrade [providers,...]",
		Short: providerUpgradeHelpMsg,
		Long:  providerUpgradeHelpMsg,
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			_, diags := c.SyncProviders(ctx, args...)
			errors.CaptureDiagnostics(diags, map[string]string{"command": "provider_upgrade"})
			if diags.HasErrors() {
				return fmt.Errorf("failed to sync providers")
			}
			return nil
		}),
	}

	providerDowngradeHelpMsg = "Downgrades one or more providers schema version based on config.hcl"
	providerDowngradeCmd     = &cobra.Command{
		Use:   "downgrade [providers,...]",
		Short: providerDowngradeHelpMsg,
		Long:  providerDowngradeHelpMsg,
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			_, diags := c.SyncProviders(ctx, args...)
			errors.CaptureDiagnostics(diags, map[string]string{"command": "provider_downgrade"})
			if diags.HasErrors() {
				return fmt.Errorf("failed to sync providers")
			}
			return nil
		}),
	}

	providerForce       bool
	providerDropHelpMsg = "Drops provider schema from database"
	providerDropCmd     = &cobra.Command{
		Use:   "drop [provider]",
		Short: providerDropHelpMsg,
		Long:  providerDropHelpMsg,
		Args:  cobra.ExactArgs(1),
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			if !providerForce {
				ui.ColorizedOutput(ui.ColorWarning, "WARNING! This will drop all tables for the given provider. If you wish to continue, use the --force flag.\n")
				return diag.FromError(fmt.Errorf("if you wish to continue, use the --force flag"), diag.USER)
			}
			diags := c.DropProvider(ctx, args[0])
			errors.CaptureDiagnostics(diags, map[string]string{"command": "provider_drop"})
			if diags.HasErrors() {
				return fmt.Errorf("failed to drop provider %s", args[0])
			}
			return nil
		}),
	}

	providerBuildSchemaHelpMsg = "Builds provider schema on database"
	providerBuildSchemaCmd     = &cobra.Command{
		Use:   "build-schema [provider]",
		Short: providerBuildSchemaHelpMsg,
		Long:  providerBuildSchemaHelpMsg,
		Args:  cobra.MaximumNArgs(1),
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			_, diags := c.SyncProviders(ctx, args...)
			errors.CaptureDiagnostics(diags, map[string]string{"command": "provider_build_schema"})
			if diags.HasErrors() {
				return fmt.Errorf("failed to sync providers")
			}
			return nil
		}),
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
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			diags := c.DownloadProviders(ctx)
			errors.CaptureDiagnostics(diags, map[string]string{"command": "provider_download"})
			if diags.HasErrors() {
				return fmt.Errorf("failed to download providers")
			}
			return nil
		}),
	}

	lastUpdate                 time.Duration
	dryRun                     bool
	providerRemoveStaleHelpMsg = "Remove stale resources from one or more providers in database"
	providerRemoveStaleCmd     = &cobra.Command{
		Use:   "purge [provider]",
		Short: providerRemoveStaleHelpMsg,
		Long:  providerRemoveStaleHelpMsg,
		Args:  cobra.MinimumNArgs(1),
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			diags := c.RemoveStaleData(ctx, lastUpdate, dryRun, args)
			errors.CaptureDiagnostics(diags, map[string]string{"command": "provider_purge"})
			if diags.HasErrors() {
				return fmt.Errorf("failed to remove stale data")
			}
			return nil
		}),
	}
)

func init() {
	providerRemoveStaleCmd.Flags().DurationVar(&lastUpdate, "last-update", time.Hour*1,
		"last-update is the duration from current time we want to remove resources from the database. "+
			"For example 24h will remove all resources that were not update in last 24 hours. Duration is a string with optional unit suffix such as \"2h45m\" or \"7d\"")
	providerRemoveStaleCmd.Flags().BoolVar(&dryRun, "dry-run", true, "")
	providerDropCmd.Flags().BoolVar(&providerForce, "force", false, "Really drop tables for the provider")
	providerCmd.AddCommand(providerDownloadCmd, providerUpgradeCmd, providerDowngradeCmd, providerDropCmd, providerBuildSchemaCmd, providerRemoveStaleCmd)
	providerCmd.SetUsageTemplate(usageTemplateWithFlags)
	rootCmd.AddCommand(providerCmd)
}
