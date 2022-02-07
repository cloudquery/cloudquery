package cmd

import (
	"context"

	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
)

const policyDescribeHelpMsg = `Describe CloudQuery policy`

var (
	describePolicyCmd = &cobra.Command{
		Use:   "describe",
		Short: policyDescribeHelpMsg,
		Long:  policyDescribeHelpMsg,
		Example: `
  # Describe official policy
  cloudquery policy describe aws
  
  # The following will be the same as above
  # Official policies are hosted here: https://github.com/cloudquery-policies
  cloudquery policy describe aws//cis-1.2.0
	
  # Describe community policy
  cloudquery policy describe github.com/COMMUNITY_GITHUB_ORG/aws

  # See https://hub.cloudquery.io for additional policies.`,
		Args: cobra.ExactArgs(1),
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			return c.DescribePolicies(ctx, args[0])
		}),
	}
)

func init() {
	describePolicyCmd.SetUsageTemplate(usageTemplateWithFlags)
	policyCmd.AddCommand(describePolicyCmd)
}
