package cmd

import (
	"context"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/internal/signalcontext"
	"github.com/cloudquery/cloudquery/pkg/module/drift"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const driftModuleID = "drift"

var (
	driftCmd = &cobra.Command{
		Use:   "drift",
		Short: "Drift Module",
		Long:  "Drift Module",
		Args:  cobra.MinimumNArgs(1),
	}

	driftInitCmd = &cobra.Command{
		Use:   "init",
		Short: "Generate config for drift",
		Long:  "Generate config for drift",
		RunE: func(cmd *cobra.Command, args []string) error {
			configPath := viper.GetString("configPath")
			ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
			c, err := console.CreateClient(ctx, configPath)
			if err != nil {
				return err
			}
			defer c.Client().Close()
			c.GenModuleConfig(ctx, driftModuleID)
			return nil
		},
	}

	driftScanCmd = &cobra.Command{
		Use:   "scan",
		Short: "Scan for drifts",
		Long:  "Scan for drifts between cloud provider and IaC",
		RunE: func(cmd *cobra.Command, args []string) error {
			configPath := viper.GetString("configPath")
			ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
			c, err := console.CreateClient(ctx, configPath)
			if err != nil {
				return err
			}
			defer c.Client().Close()

			return c.CallModule(ctx, console.ModuleCallRequest{
				Name:          driftModuleID,
				Params:        driftParams,
				ModConfigPath: driftConfigPath,
				OutputPath:    driftOutputPath,
			})
		},
	}

	driftParams drift.RunParams

	driftOutputPath, driftConfigPath string
)

func init() {
	flags := driftScanCmd.Flags()

	// generic flags
	flags.StringVar(&driftOutputPath, "output", "", "Generate a new file at the given path with the output")
	flags.StringVar(&driftConfigPath, "drift-config", getDefaultModuleConfigFile(driftModuleID), "Use the given drift config file")

	// flags handled by the drift package
	flags.BoolVar(&driftParams.Debug, "debug", false, "Show debug output")
	flags.StringSliceVar(&driftParams.AccountIDs, "account-ids", nil, "Use only specified cloud account IDs")
	flags.StringSliceVar(&driftParams.TfBackendNames, "tf-backend-names", nil, "Filter by Terraform backend names")
	flags.StringVar(&driftParams.TfMode, "tf-mode", "managed", "Set Terraform mode")
	flags.StringVar(&driftParams.TfProvider, "tf-provider", "", "Set Terraform provider (defaults to cloud provider name)")
	flags.BoolVar(&driftParams.ForceDeep, "deep", false, "Force deep mode")
	flags.BoolVar(&driftParams.ListManaged, "list-managed", false, "List managed resources in output")

	driftCmd.SetUsageTemplate(usageTemplateWithFlags)
	driftCmd.AddCommand(driftScanCmd)
	driftCmd.AddCommand(driftInitCmd)
	rootCmd.AddCommand(driftCmd)
}

func getDefaultModuleConfigFile(modName string) string {
	proposedFilename := modName + ".hcl"
	fs := afero.NewOsFs()
	i, err := fs.Stat(proposedFilename)
	if err != nil {
		return ""
	}
	if i.IsDir() {
		return ""
	}

	return proposedFilename
}
