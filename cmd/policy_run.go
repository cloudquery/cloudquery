package cmd

import (
	"context"

	"github.com/cloudquery/cloudquery/pkg/ui/console"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const policyRunHelpMsg = "Executes a policy on CloudQuery database"

var (
	policyRunCmd = &cobra.Command{
		Use:   "run",
		Short: policyRunHelpMsg,
		Long:  policyRunHelpMsg,
		Example: `
  # Download & Run the policies defined in your config
  cloudquery policy run

  # Run a specific policy that is not defined in the config.hcl
  # Run official policy
  cloudquery policy run aws

  # The following will be the same as above
  # Official policies are hosted here: https://github.com/cloudquery-policies
  cloudquery policy run aws//cis_v1.2.0
	
  # Run community policy
  cloudquery policy run github.com/COMMUNITY_GITHUB_ORG/aws

  # See https://hub.cloudquery.io for additional policies.`,
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				return c.RunPolicies(ctx, args[0], outputDir, noResults)
			}
			return c.RunPolicies(ctx, "", outputDir, noResults)
		}),
		Args: cobra.MaximumNArgs(1),
	}
	outputDir string
	noResults bool
)

func init() {
	flags := policyRunCmd.Flags()
	flags.StringVar(&outputDir, "output-dir", "", "Generates a new file for each policy at the given dir with the output")
	flags.BoolVar(&noResults, "no-results", false, "Do not show policies results")
	flags.Bool("disable-fetch-check", false, "Disable checking if a respective fetch happened before running policies")

	_ = viper.BindPFlag("disable-fetch-check", flags.Lookup("disable-fetch-check"))

	policyRunCmd.SetUsageTemplate(usageTemplateWithFlags)
	policyCmd.AddCommand(policyRunCmd)
}
