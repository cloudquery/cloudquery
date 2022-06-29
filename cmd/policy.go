package cmd

import (
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/errors"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	policyHelpMsg         = `Download and run CloudQuery policy`
	policyDescribeHelpMsg = `Describe CloudQuery policy`
	policyRunHelpMsg      = "Executes a policy on CloudQuery database"
	policySnapshotHelpMsg = `Take database snapshot of all tables included in a CloudQuery policy`
	policyTestHelpMsg     = "Tests policy against a precompiled set of database snapshots"
	policyValidateHelpMsg = "Validate policy for any issues and diagnostics"
	policyPruneHelpMsg    = "Prune policy executions from the database which are older than the relative time specified"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), getConfigFile(), true, nil, instanceId)
			if err != nil {
				return err
			}
			return c.DescribePolicies(cmd.Context(), args[0])
		},
	}

	policyDownloadCmd = &cobra.Command{
		Use:        "download",
		Deprecated: "Use 'policy run' command directly instead. If you need to download a policy to your machine, you can use 'git clone <URL FOR GIT REPO>'. See https://docs.cloudquery.io/docs/cli/policy/sources for more information on how to use the `policy run` command ",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("'policy download' command has been deprecated. Use the 'policy run' command directly instead")
		},
	}

	outputDir    string
	noResults    bool
	storeResults bool
	policyRunCmd = &cobra.Command{
		Use:   "run <policy>",
		Short: policyRunHelpMsg,
		Long:  policyRunHelpMsg,
		Example: `
	# Run an official policy
	# Official policies are available on our hub: https://hub.cloudquery.io/policies
	cloudquery policy run aws

	# Run a sub-policy of an official policy
	cloudquery policy run aws//cis_v1.2.0

	# Run a policy from a GitHub repository
	cloudquery policy run github.com/<repo-owner>/<repo-name>

	# Run a policy from a local directory
	cloudquery policy run ./PATH_TO_POLICY_DIRECTORY/

	# See https://hub.cloudquery.io for additional policies.
	# See https://docs.cloudquery.io/docs/tutorials/policies/policies-overview for instructions on writing policies.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			source := args[0]
			c, err := console.CreateClient(cmd.Context(), getConfigFile(), false, nil, instanceId)
			if err != nil {
				return err
			}
			diags := c.RunPolicies(cmd.Context(), source, outputDir, noResults, storeResults)
			errors.CaptureDiagnostics(diags, map[string]string{"command": "policy_run"})
			if diags.HasErrors() {
				return fmt.Errorf("policy has one or more errors, check logs")
			}
			return nil
		},
		Args: cobra.ExactArgs(1),
	}

	policyTestCmd = &cobra.Command{
		Use:   "test",
		Short: policyTestHelpMsg,
		Long:  policyTestHelpMsg,
		Example: `
  # Download & Run the policies defined in your config
  cloudquery policy test path/to/policy.hcl path/to/snapshot/dir selector
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), getConfigFile(), false, nil, instanceId)
			if err != nil {
				return err
			}
			err = c.TestPolicies(cmd.Context(), args[0], args[1])
			errors.CaptureError(err, map[string]string{"command": "policy_test"})
			return err
		},
		Args: cobra.ExactArgs(2),
	}

	snapshotPolicyCmd = &cobra.Command{
		Use:   "snapshot",
		Short: policySnapshotHelpMsg,
		Long:  policySnapshotHelpMsg,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), getConfigFile(), false, nil, instanceId)
			if err != nil {
				return err
			}
			err = c.SnapshotPolicy(cmd.Context(), args[0], args[1])
			errors.CaptureError(err, map[string]string{"command": "policy_snapshot"})
			return err
		},
	}

	validatePolicyCmd = &cobra.Command{
		Use:   "validate",
		Short: policyValidateHelpMsg,
		Long:  policyValidateHelpMsg,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), getConfigFile(), false, nil, instanceId)
			if err != nil {
				return err
			}
			diags := c.ValidatePolicy(cmd.Context(), args[0])
			errors.CaptureDiagnostics(diags, map[string]string{"command": "policy_validate"})
			if diags.HasErrors() {
				return fmt.Errorf("policy validate has one or more errors, check logs")
			}
			return nil
		},
	}

	prunePolicyCmd = &cobra.Command{
		Use:   "prune",
		Short: policyPruneHelpMsg,
		Long:  policyPruneHelpMsg,
		Example: `
  # Prune the policy executions which are older than the relative time specified
  cloudquery policy prune 24h`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), getConfigFile(), false, nil, instanceId)
			if err != nil {
				return err
			}
			retentionPeriod := args[0]
			diags := c.PrunePolicyExecutions(cmd.Context(), retentionPeriod)
			errors.CaptureDiagnostics(diags, map[string]string{"command": "policy_prune"})
			if diags.HasErrors() {
				return fmt.Errorf("policy prune has one or more errors, check logs")
			}
			return nil
		},
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
	flags.BoolVar(&storeResults, "enable-db-persistence", false, "Enable storage of policy output in database")
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

	policyCmd.AddCommand(validatePolicyCmd)

	policyCmd.AddCommand(prunePolicyCmd)

	policyCmd.SetUsageTemplate(usageTemplateWithFlags)
	rootCmd.AddCommand(policyCmd)
}
