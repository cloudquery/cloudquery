package policy

import (
	"github.com/cloudquery/cloudquery/cmd/utils"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
)

const (
	describeShort   = "Describe CloudQuery policy"
	describeExample = `
# Describe official policy
cloudquery policy describe aws

# The following will be the same as above
# Official policies are hosted here: https://github.com/cloudquery-policies
cloudquery policy describe aws//cis-1.2.0

# Describe community policy
cloudquery policy describe github.com/COMMUNITY_GITHUB_ORG/aws

# See https://hub.cloudquery.io for additional policies.
`
)

func newCmdPolicyDescribe() *cobra.Command {
	describePolicyCmd := &cobra.Command{
		Use:     "describe",
		Short:   describeShort,
		Long:    describeShort,
		Example: describeExample,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), utils.GetConfigFile(), true, nil, utils.InstanceId)
			if err != nil {
				return err
			}
			return c.DescribePolicies(cmd.Context(), args[0])
		},
	}
	return describePolicyCmd
}
