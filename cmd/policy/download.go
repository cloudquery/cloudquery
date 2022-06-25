package policy

import (
	"github.com/cloudquery/cloudquery/cmd/util"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
)

const (
	downloadShort   = "Download a policy from the CloudQuery Policy Hub"
	downloadExample = `
  # Download official policy
  cloudquery policy download aws
  
  # The following will be the same as above
  # Official policies are hosted here: https://github.com/cloudquery-policies
  cloudquery policy download aws//cis-1.2.0
	
  # Download community policy
  cloudquery policy download github.com/COMMUNITY_GITHUB_ORG/aws

  # See https://hub.cloudquery.io for additional policies.`
)

func NewCmdPolicyDownload(parent policyOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "download GITHUB_REPO",
		Short:   downloadShort,
		Long:    downloadShort,
		Example: downloadExample,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), util.GetConfigFile(parent.Config), true, nil, util.InstanceId)
			if err != nil {
				return err
			}
			return c.DownloadPolicy(cmd.Context(), args)
		},
	}
	return cmd
}
