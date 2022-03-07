package cmd

import (
	"context"

	"github.com/cloudquery/cloudquery/pkg/ui/console"

	"github.com/spf13/cobra"
)

const policyTestHelpMsg = "Tests policy against a precompiled set of database snapshots"

var (
	policyTestCmd = &cobra.Command{
		Use:   "test",
		Short: policyTestHelpMsg,
		Long:  policyTestHelpMsg,
		Example: `
  # Download & Run the policies defined in your config
  cloudquery policy test path/to/policy.hcl path/to/snapshot/dir selector
	`,
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			return c.TestPolicies(ctx, args[0], args[1])
		}),
		Args: cobra.ExactArgs(2),
	}
)

func init() {
	flags := policyTestCmd.Flags()
	flags.StringVar(&outputDir, "output-dir", "", "Generates a new file for each policy at the given dir with the output")
	flags.BoolVar(&noResults, "no-results", false, "Do not show policies results")
	policyRunCmd.SetUsageTemplate(usageTemplateWithFlags)
	policyCmd.AddCommand(policyTestCmd)
}
