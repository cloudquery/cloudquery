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
	policyDownloadHelpMsg = "Download a policy from the CloudQuery Policy Hub"
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
			cfgPath := viper.GetString("configPath")
			c, err := console.CreateClient(cmd.Context(), cfgPath, true, nil, instanceId)
			if err != nil {
				return err
			}
			return c.DescribePolicies(cmd.Context(), args[0])
		},
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
		RunE: func(cmd *cobra.Command, args []string) error {
			cfgPath := viper.GetString("configPath")
			c, err := console.CreateClient(cmd.Context(), cfgPath, true, nil, instanceId)
			if err != nil {
				return err
			}
			return c.DownloadPolicy(cmd.Context(), args)
		},
	}

	outputDir    string
	noResults    bool
	storeResults bool
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
		RunE: func(cmd *cobra.Command, args []string) error {
			source := args[0]
			cfgPath := viper.GetString("configPath")
			c, err := console.CreateClient(cmd.Context(), cfgPath, true, nil, instanceId)
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
			cfgPath := viper.GetString("configPath")
			c, err := console.CreateClient(cmd.Context(), cfgPath, true, nil, instanceId)
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
			cfgPath := viper.GetString("configPath")
			c, err := console.CreateClient(cmd.Context(), cfgPath, true, nil, instanceId)
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
			cfgPath := viper.GetString("configPath")
			c, err := console.CreateClient(cmd.Context(), cfgPath, true, nil, instanceId)
			if err != nil {
				return err
			}
			diags := c.ValidatePolicy(cmd.Context(), args[0])
			errors.CaptureDiagnostics(diags, map[string]string{"command": "policy_validate"})
			return fmt.Errorf("policy validate has one or more errors, check logs")
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
			cfgPath := viper.GetString("configPath")
			c, err := console.CreateClient(cmd.Context(), cfgPath, true, nil, instanceId)
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
