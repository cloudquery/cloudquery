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

const moduleRunHelpMsg = `Executes a CloudQuery module`

var (
	moduleRunCmd = &cobra.Command{
		Use:   "run MODULE_NAME [OPTIONS]",
		Short: moduleRunHelpMsg,
		Long:  moduleRunHelpMsg,
		//Example: "",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			configPath := viper.GetString("configPath")
			ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
			c, err := console.CreateClient(ctx, configPath)
			if err != nil {
				return err
			}
			defer c.Client().Close()
			return c.RunModule(ctx, args, moduleOutputPath, moduleConfigPath)
		},
	}
	moduleOutputPath, moduleConfigPath string
)

func init() {
	flags := moduleRunCmd.Flags()
	flags.StringVar(&moduleOutputPath, "output", "", "Generates a new file at the given path with the output")
	flags.StringVar(&moduleConfigPath, "modconfig", "", "Use the given module config file")
	moduleRunCmd.SetUsageTemplate(usageTemplateWithFlags)
	moduleCmd.AddCommand(moduleRunCmd)
}
