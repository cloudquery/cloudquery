package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/internal/signalcontext"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
)

const policyRunHelpMsg = "Executes a policy on CloudQuery database"

var (
	policyRunCmd = &cobra.Command{
		Use:   "run",
		Short: policyRunHelpMsg,
		Long:  policyRunHelpMsg,
		Example: `
  # Download & Run the policies that defined in the config.hcl
  cloudquery policy run

  # Run a specific policy by it's name
  cloudquery policy run --policy my_aws_policy

  # See https://hub.cloudquery.io for additional policies.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			configPath := viper.GetString("configPath")
			ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
			c, err := console.CreateClient(ctx, configPath)
			if err != nil {
				return err
			}
			return c.RunPolicies(ctx, policyName, outputDir, stopOnFailure, skipVersioning, failOnViolation, noResults)
		},
	}
	outputDir       string
	stopOnFailure   bool
	policyName      string
	skipVersioning  bool
	failOnViolation bool
	noResults       bool
)

func init() {
	flags := policyRunCmd.Flags()
	flags.StringVar(&policyName, "policy", "", "Select specific policy to run")
	flags.StringVar(&outputDir, "output", "", "Generates a new file for each policy at the given dir with the output")
	flags.BoolVar(&stopOnFailure, "stop-on-failure", false, "Stops the policy execution on the first failure")
	flags.BoolVar(&failOnViolation, "fail-on-violation", false, "Return non zero exit code if one of the policy is violated")
	flags.BoolVar(&skipVersioning, "skip-versioning", false, "Skip policy versioning and use latest files")
	flags.BoolVar(&noResults, "no-results", false, "Do not show policies results")
	policyRunCmd.SetUsageTemplate(usageTemplateWithFlags)
	policyCmd.AddCommand(policyRunCmd)
}
