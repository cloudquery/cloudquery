package cmd

import (
	"github.com/cloudquery/cloudquery/internal/analytics"
	"github.com/cloudquery/cloudquery/pkg/errors"
	"github.com/cloudquery/cloudquery/pkg/module/drift"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			cfgPath := viper.GetString("configPath")
			cfgMutator := filterConfigProviders(args)
			c, err := console.CreateClient(cmd.Context(), cfgPath, false, cfgMutator, instanceId)
			if err != nil {
				return err
			}
			driftParams.StateFiles = args
			diags := c.CallModule(cmd.Context(), console.ModuleCallRequest{
				Name:       "drift",
				Params:     driftParams,
				Profile:    driftProfile,
				OutputPath: driftOutputPath,
			})
			analytics.Capture("drift", c.Providers, nil, diags)
			errors.CaptureError(diags, map[string]string{"command": "drift"})
			return diags
		},
	}

	driftParams     drift.RunParams
	driftProfile    string
	driftOutputPath string
)

func init() {
	flags := driftScanCmd.Flags()

	// generic flags
	flags.StringVar(&driftOutputPath, "output", "", "Generate a new file at the given path with the output")
	flags.StringVar(&driftProfile, "profile", "", "Specify drift profile")

	// flags handled by the drift package
	flags.BoolVar(&driftParams.Debug, "debug", false, "Show debug output")
	flags.BoolVar(&driftParams.ForceDeep, "deep", false, "Force deep mode")
	flags.BoolVar(&driftParams.ListManaged, "list-managed", false, "List managed resources in output")

	driftCmd.SetUsageTemplate(usageTemplateWithFlags)
	driftCmd.AddCommand(driftScanCmd)
	rootCmd.AddCommand(driftCmd)
}
