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
		Args:  cobra.ExactArgs(1),
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			_ = c.DescribePolicies(ctx, args, args[0], skipVersioning)
			return nil
		}),
	}
)

func init() {
	describePolicyCmd.SetUsageTemplate(usageTemplateWithFlags)
	describePolicyCmd.Flags().BoolVar(&skipVersioning, "skip-versioning", false, "Skip policy versioning and use latest files")
	policyCmd.AddCommand(describePolicyCmd)
}
