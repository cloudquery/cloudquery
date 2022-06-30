package policy

import (
	"github.com/cloudquery/cloudquery/cmd/utils"
	"github.com/cloudquery/cloudquery/pkg/errors"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
)

const (
	testShort   = "Tests policy against a precompiled set of database snapshots"
	testExample = `
	# Download & Run the policies defined in your config
	cloudquery policy test path/to/policy.hcl path/to/snapshot/dir selector
		`
)

func newCmdPolicyTest() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "test",
		Short:   testShort,
		Long:    testShort,
		Example: testExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), utils.GetConfigFile(), false, nil, utils.InstanceId)
			if err != nil {
				return err
			}
			err = c.TestPolicies(cmd.Context(), args[0], args[1])
			errors.CaptureError(err, map[string]string{"command": "policy_test"})
			return err
		},
		Args: cobra.ExactArgs(2),
	}
	flags := cmd.Flags()
	flags.StringVar(&outputDir, "output-dir", "", "Generates a new file for each policy at the given dir with the output")
	flags.BoolVar(&noResults, "no-results", false, "Do not show policies results")
	return cmd
}
