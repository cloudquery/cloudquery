package cmd

import (
	"context"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/internal/signalcontext"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const moduleHelpMsg = `CloudQuery module commands`

var (
	moduleCmd = &cobra.Command{
		Use:   "module MODNAME [SUBCOMMAND] [ARGS]",
		Short: moduleHelpMsg,
		Long:  moduleHelpMsg,
		Args:  cobra.MinimumNArgs(1),
	}
	moduleOutputPath, moduleConfigPath string

	moduleGenCmd = &cobra.Command{
		Use:   "module-gen MODNAME [SUBCOMMAND] [ARGS]",
		Short: "Generate config for the given module",
		Long:  "Generate config for the given module",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			configPath := viper.GetString("configPath")
			ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
			c, err := console.CreateClient(ctx, configPath)
			if err != nil {
				return err
			}
			defer c.Client().Close()
			c.GenModuleConfig(ctx, args)
			return nil
		},
	}
)

func init() {
	moduleCmd.SetUsageTemplate(usageTemplateWithFlags)
	flags := moduleCmd.PersistentFlags()
	flags.StringVar(&moduleOutputPath, "output", "", "Generates a new file at the given path with the output")
	flags.StringVar(&moduleConfigPath, "modconfig", "", "Use the given module config file")
	rootCmd.AddCommand(moduleCmd)

	moduleGenCmd.SetUsageTemplate(usageTemplateWithFlags)
	rootCmd.AddCommand(moduleGenCmd)
}
