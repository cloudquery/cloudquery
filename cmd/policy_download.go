package cmd

import (
	"context"

	"github.com/cloudquery/cloudquery/pkg/ui/console"

	"github.com/spf13/cobra"
)

const policyDownloadHelpMsg = "Download a policy from the CloudQuery Policy Hub"

var (
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
)

func init() {
	policyDownloadCmd.SetUsageTemplate(usageTemplateWithFlags)
	policyCmd.AddCommand(policyDownloadCmd)
}
