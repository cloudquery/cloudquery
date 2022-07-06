package policy

import (
	"fmt"

	"github.com/cloudquery/cloudquery/cmd/utils"
	"github.com/cloudquery/cloudquery/pkg/errors"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	outputDir    string
	noResults    bool
	storeResults bool
)

const (
	runShort     = "Executes a policy on CloudQuery database"
	exampleShort = `
# Run an official policy
# Official policies are available on our hub: https://hub.cloudquery.io/policies
cloudquery policy run aws

# Run a sub-policy of an official policy
cloudquery policy run aws//cis_v1.2.0

# Run a policy from a GitHub repository
cloudquery policy run github.com/<repo-owner>/<repo-name>

# Run a policy from a local directory
cloudquery policy run ./<path-to-local-directory>

# See https://hub.cloudquery.io for additional policies
# See https://docs.cloudquery.io/docs/tutorials/policies/policies-overview for instructions on writing policies`
)

func newCmdPolicyRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "run <policy>",
		Short:   runShort,
		Long:    runShort,
		Example: exampleShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			source := args[0]
			c, err := console.CreateClient(cmd.Context(), utils.GetConfigFile(), false, nil, utils.InstanceId)
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
	flags := cmd.Flags()
	flags.StringVar(&outputDir, "output-dir", "", "Generates a new file for each policy at the given dir with the output")
	flags.BoolVar(&noResults, "no-results", false, "Do not show policies results")
	flags.BoolVar(&storeResults, "enable-db-persistence", false, "Enable storage of policy output in database")
	flags.Bool("disable-fetch-check", false, "Disable checking if a respective fetch happened before running policies")
	_ = viper.BindPFlag("disable-fetch-check", flags.Lookup("disable-fetch-check"))
	return cmd
}
