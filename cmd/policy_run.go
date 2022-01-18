package cmd

import (
	"context"

	"github.com/cloudquery/cloudquery/pkg/ui/console"

	"github.com/spf13/cobra"
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

  # Run a specific policy by it's name
  cloudquery policy run --policy my_aws_policy

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
	policyRunCmd.SetUsageTemplate(usageTemplateWithFlags)
	policyCmd.AddCommand(policyRunCmd)
}
