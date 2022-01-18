package cmd

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/client"
	"github.com/cloudquery/cloudquery/pkg/ui/console"

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
  ./cloudquery provider download
  # Upgrades all providers specified in config.hcl
  ./cloudquery provider upgrade 
  # Upgrade one or more providers
  ./cloudquery provider upgrade aws
  # Downgrades all providers specified in config.hcl
  ./cloudquery provider downgrade 
  # Downgrades one or more providers
  ./cloudquery provider downgrade aws, gcp
  # Drop provider schema, running fetch again will recreate all tables unless --skip-build-tables is specified
  ./cloudquery provider drop aws
  # build provider schema
  ./cloudquery provider build-schema aws
`,
		Version: client.Version,
	}

	providerUpgradeHelpMsg = "Upgrades one or more providers schema version based on config.hcl"
	providerUpgradeCmd     = &cobra.Command{
		Use:   "upgrade [providers,...]",
		Short: providerUpgradeHelpMsg,
		Long:  providerUpgradeHelpMsg,
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			return c.UpgradeProviders(ctx, args)
		}),
	}

	providerDowngradeHelpMsg = "Downgrades one or more providers schema version based on config.hcl"
	providerDowngradeCmd     = &cobra.Command{
		Use:   "downgrade [providers,...]",
		Short: providerDowngradeHelpMsg,
		Long:  providerDowngradeHelpMsg,
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			return c.DowngradeProviders(ctx, args)
		}),
	}

	providerDropHelpMsg = "Drops provider schema from database"
	providerDropCmd     = &cobra.Command{
		Use:   "drop [provider]",
		Short: providerDropHelpMsg,
		Long:  providerDropHelpMsg,
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("missing provider name")
			}
			_ = c.DropProvider(ctx, args[0])
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
			if len(args) == 1 {
				return c.BuildProviderTables(ctx, args[0])
			}
			return c.BuildAllProviderTables(ctx)
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
			return c.DownloadProviders(ctx)
		}),
	}
)

func init() {
	providerCmd.AddCommand(providerDownloadCmd, providerUpgradeCmd, providerDowngradeCmd, providerDropCmd, providerBuildSchemaCmd)
	rootCmd.AddCommand(providerCmd)
}
