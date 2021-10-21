package cmd

import (
	"context"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/internal/signalcontext"
	"github.com/cloudquery/cloudquery/pkg/module/drift"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	driftCmd = &cobra.Command{
		Use:   "drift",
		Short: "Drift Module",
		Long:  "Drift Module",
		Args:  cobra.MinimumNArgs(1),
	}

	driftRunCmd = &cobra.Command{
		Use:   "run",
		Short: "Detect drifts",
		Long:  "Detect drifts between cloud provider and IaC",
		RunE: func(cmd *cobra.Command, args []string) error {
			configPath := viper.GetString("configPath")
			ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
			c, err := console.CreateClient(ctx, configPath)
			if err != nil {
				return err
			}
			defer c.Client().Close()

			return c.CallModule(ctx, "drift", moduleOutputPath, moduleConfigPath, driftParams)
		},
	}

	driftParams drift.RunParams
)

func init() {
	driftCmd.PersistentFlags().BoolVar(&driftParams.Debug, "debug", false, "Show debug output")

	flags := driftRunCmd.Flags()
	flags.StringVar(&driftParams.TfBackendName, "tf-backend-name", "mylocal", "Set Terraform backend name")
	flags.StringVar(&driftParams.TfMode, "tf-mode", "managed", "Set Terraform mode")
	flags.StringVar(&driftParams.TfProvider, "tf-provider", "", "Set Terraform provider (defaults to cloud provider name)")
	flags.BoolVar(&driftParams.ForceDeep, "deep", false, "Force deep mode")
	driftRunCmd.SetUsageTemplate(usageTemplateWithFlags)
	driftCmd.AddCommand(driftRunCmd)
	moduleCmd.AddCommand(driftCmd)
}
