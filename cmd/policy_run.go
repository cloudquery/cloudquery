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

const policyRunHelpMsg = "Runs a policy"

var (
	policyRunCmd = &cobra.Command{
		Use:   "run GITHUB_REPO [PATH_IN_REPO]",
		Short: policyRunHelpMsg,
		Long:  policyRunHelpMsg,
		Example: `
  # Download & Run official policy from Policy Hub
  cloudquery policy run aws-cis-1.2.0

  # Download & Run official policy from Policy Hub (equivalent to the above)
  # Official policies hosted at https://github.com/cloudquery-policies
  cloudquery policy run cloudquery-policies/aws-cis-1.2.0

  # Run without downloading
  cloudquery policy download aws-cis-1.2.0
  cloudquery policy run --skip-download aws-cis-1.2.0

  # See https://hub.cloudquery.io for additional policies.`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			configPath := viper.GetString("configPath")
			ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
			c, err := console.CreateClient(ctx, configPath)
			if err != nil {
				return err
			}
			defer c.Client().Close()
			err = c.DownloadPolicy(ctx, args)
			if err != nil {
				return err
			}
			_ = c.RunPolicy(ctx, args, subPath, outputPath, stopOnFailure, skipVersioning)
			return nil
		},
	}
	skipDownload   bool
	subPath        string
	outputPath     string
	stopOnFailure  bool
	skipVersioning bool
)

func init() {
	flags := policyRunCmd.Flags()
	flags.BoolVar(&skipDownload, "skip-download", false, "Skip downloading the policy repository")
	flags.StringVar(&subPath, "sub-path", "", "Forces the policy run command to only execute this sub policy/query")
	flags.StringVar(&outputPath, "output", "", "Generates a new file at the given path with the output")
	flags.BoolVar(&stopOnFailure, "stop-on-failure", false, "Stops the execution on the first failure")
	flags.BoolVar(&skipVersioning, "skip-versioning", false, "Skip policy versioning and use latest files")
	policyRunCmd.SetUsageTemplate(usageTemplateWithFlags)
	policyCmd.AddCommand(policyRunCmd)
}
