package cmd

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/errors"

	"github.com/spf13/viper"

	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
)

const (
	policyHelpMsg         = `Download and run CloudQuery policy`
	policyDescribeHelpMsg = `Describe CloudQuery policy`
	policyDownloadHelpMsg = "Download a policy from the CloudQuery Policy Hub"
	policyRunHelpMsg      = "Executes a policy on CloudQuery database"
	policySnapshotHelpMsg = `Take database snapshot of all tables included in a CloudQuery policy`
	policyTestHelpMsg     = "Tests policy against a precompiled set of database snapshots"
)

var (
	policyCmd = &cobra.Command{
		Use:   "policy SUBCOMMAND",
		Short: policyHelpMsg,
		Long:  policyHelpMsg,
	}
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

	policyDownloadCmd = &cobra.Command{
		Use:   "download GITHUB_REPO",
		Short: policyDownloadHelpMsg,
		Long:  policyDownloadHelpMsg,
		Example: `
  # Download official policy
  cloudquery policy download aws
  
  # The following will be the same as above
  # Official policies are hosted here: https://github.com/cloudquery-policies
  cloudquery policy download aws//cis-1.2.0
	
  # Download community policy
  cloudquery policy download github.com/COMMUNITY_GITHUB_ORG/aws

  # See https://hub.cloudquery.io for additional policies.`,
		Args: cobra.ExactArgs(1),
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			_ = c.DownloadPolicy(ctx, args)
			return nil
		}),
	}

	outputDir    string
	noResults    bool
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
			source := ""
			if len(args) == 1 {
				source = args[0]
			}
			diags := c.RunPolicies(ctx, source, outputDir, noResults)
			errors.CaptureDiagnostics(diags, map[string]string{"command": "policy_run"})
			if diags.HasErrors() {
				return fmt.Errorf("provider has one or more errors, check logs")
			}
			return nil
		}),
		Args: cobra.MaximumNArgs(1),
	}

	policyTestCmd = &cobra.Command{
		Use:   "test",
		Short: policyTestHelpMsg,
		Long:  policyTestHelpMsg,
		Example: `
  # Download & Run the policies defined in your config
  cloudquery policy test path/to/policy.hcl path/to/snapshot/dir selector
	`,
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			err := c.TestPolicies(ctx, args[0], args[1])
			errors.CaptureError(err, map[string]string{"command": "policy_test"})
			return err
		}),
		Args: cobra.ExactArgs(2),
	}

	snapshotPolicyCmd = &cobra.Command{
		Use:   "snapshot",
		Short: policySnapshotHelpMsg,
		Long:  policySnapshotHelpMsg,
		Args:  cobra.ExactArgs(2),
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			err := c.SnapshotPolicy(ctx, args[0], args[1])
			errors.CaptureError(err, map[string]string{"command": "policy_snapshot"})
			return err
		}),
	}
)

func init() {
	describePolicyCmd.SetUsageTemplate(usageTemplateWithFlags)
	policyCmd.AddCommand(describePolicyCmd)
	policyDownloadCmd.SetUsageTemplate(usageTemplateWithFlags)
	policyCmd.AddCommand(policyDownloadCmd)

	flags := policyRunCmd.Flags()
	flags.StringVar(&outputDir, "output-dir", "", "Generates a new file for each policy at the given dir with the output")
	flags.BoolVar(&noResults, "no-results", false, "Do not show policies results")
	flags.Bool("disable-fetch-check", false, "Disable checking if a respective fetch happened before running policies")
	_ = viper.BindPFlag("disable-fetch-check", flags.Lookup("disable-fetch-check"))
	policyRunCmd.SetUsageTemplate(usageTemplateWithFlags)
	policyCmd.AddCommand(policyRunCmd)

	snapshotPolicyCmd.SetUsageTemplate(usageTemplateWithFlags)
	policyCmd.AddCommand(snapshotPolicyCmd)

	flags = policyTestCmd.Flags()
	flags.StringVar(&outputDir, "output-dir", "", "Generates a new file for each policy at the given dir with the output")
	flags.BoolVar(&noResults, "no-results", false, "Do not show policies results")
	policyRunCmd.SetUsageTemplate(usageTemplateWithFlags)
	policyCmd.AddCommand(policyTestCmd)

	policyCmd.SetUsageTemplate(usageTemplateWithFlags)
	rootCmd.AddCommand(policyCmd)
}
