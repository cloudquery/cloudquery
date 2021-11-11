package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/internal/signalcontext"
	"github.com/cloudquery/cloudquery/pkg/module/drift"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
)

var (
	driftCmd = &cobra.Command{
		Use:   "drift",
		Short: "Drift Module",
		Long:  "Drift Module",
	}

	driftScanCmd = &cobra.Command{
		Use:   "scan [state files...]",
		Short: "Scan for drifts",
		Long:  "Scan for drifts between cloud provider and IaC",
		Run: handleError(func(cmd *cobra.Command, args []string) error {
			configPath := viper.GetString("configPath")
			ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
			c, err := console.CreateClient(ctx, configPath)
			if err != nil {
				return err
			}
			defer c.Client().Close()

			driftParams.StateFiles = args

			return c.CallModule(ctx, console.ModuleCallRequest{
				Name:       "drift",
				Params:     driftParams,
				Profile:    driftProfile,
				OutputPath: driftOutputPath,
			})
		}),
	}

	driftParams drift.RunParams

	driftProfile, driftOutputPath string
)

func init() {
	flags := driftScanCmd.Flags()

	// generic flags
	flags.StringVar(&driftOutputPath, "output", "", "Generate a new file at the given path with the output")
	flags.StringVar(&driftProfile, "profile", "", "Specify drift profile")

	// flags handled by the drift package
	flags.BoolVar(&driftParams.Debug, "debug", false, "Show debug output")
	flags.StringVar(&driftParams.TfMode, "tf-mode", "managed", "Set Terraform mode")
	flags.BoolVar(&driftParams.ForceDeep, "deep", false, "Force deep mode")
	flags.BoolVar(&driftParams.ListManaged, "list-managed", false, "List managed resources in output")

	driftCmd.SetUsageTemplate(usageTemplateWithFlags)
	driftCmd.AddCommand(driftScanCmd)
	rootCmd.AddCommand(driftCmd)
}
