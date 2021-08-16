package cmd

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/internal/signalcontext"
	"github.com/cloudquery/cloudquery/pkg/client"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
`,
		Version: client.Version,
	}

	providerUpgradeHelpMsg = "Upgrades one or more providers schema version based on config.hcl"
	providerUpgradeCmd     = &cobra.Command{
		Use:   "upgrade [provider]",
		Short: providerUpgradeHelpMsg,
		Long:  providerUpgradeHelpMsg,
		RunE: func(cmd *cobra.Command, args []string) error {
			configPath := viper.GetString("configPath")
			ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
			c, err := console.CreateClient(ctx, configPath)
			if err != nil {
				return err
			}
			defer c.Client().Close()
			_ = c.UpgradeProviders(ctx, args)
			return nil
		},
	}

	providerDowngradeHelpMsg = "Downgrades one or more providers schema version based on config.hcl"
	providerDowngradeCmd     = &cobra.Command{
		Use:   "downgrade [provider]",
		Short: providerDowngradeHelpMsg,
		Long:  providerDowngradeHelpMsg,
		RunE: func(cmd *cobra.Command, args []string) error {
			configPath := viper.GetString("configPath")
			ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
			c, err := console.CreateClient(ctx, configPath)
			if err != nil {
				return err
			}
			defer c.Client().Close()
			_ = c.DowngradeProviders(ctx, args)
			return nil
		},
	}

	providerDropHelpMsg = "Drops provider schema from database"
	providerDropCmd     = &cobra.Command{
		Use:   "drop provider",
		Short: providerDropHelpMsg,
		Long:  providerDropHelpMsg,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("missing provider name")
			}
			configPath := viper.GetString("configPath")
			ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
			c, err := console.CreateClient(ctx, configPath)
			if err != nil {
				return err
			}
			defer c.Client().Close()
			_ = c.DropProvider(ctx, args[0])
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
			configPath := viper.GetString("configPath")
			ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
			c, err := console.CreateClient(ctx, configPath)
			if err != nil {
				return err
			}
			defer c.Client().Close()
			return c.DownloadProviders(ctx)
		},
	}
)

func init() {
	providerCmd.AddCommand(providerDownloadCmd, providerUpgradeCmd, providerDowngradeCmd, providerDropCmd)
	rootCmd.AddCommand(providerCmd)
}
