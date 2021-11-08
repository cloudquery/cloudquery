package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/internal/signalcontext"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
)

const policyDownloadHelpMsg = "Download a policy from the CloudQuery Policy Hub"

var (
	policyDownloadCmd = &cobra.Command{
		Use:   "download GITHUB_REPO",
		Short: policyDownloadHelpMsg,
		Long:  policyDownloadHelpMsg,
		Example: `
  # Download official policy
  cloudquery policy download aws-cis-1.2.0
  
  # The following will be the same as above
  # Official policies are hosted here: https://github.com/cloudquery-policies
  cloudquery policy download cloudquery-policies/aws-cis-1.2.0

  # Download community policy
  cloudquery policy download COMMUNITY_GITHUB_ORG/aws-cis-1.2.0

  # See https://hub.cloudquery.io for additional policies.`,
		Args: cobra.ExactArgs(1),
		Run: handleError(func(cmd *cobra.Command, args []string) error {
			configPath := viper.GetString("configPath")
			ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
			c, err := console.CreateClient(ctx, configPath)
			if err != nil {
				return err
			}
			defer c.Client().Close()
			_ = c.DownloadPolicy(ctx, args)
			return nil
		}),
	}
)

func init() {
	policyDownloadCmd.SetUsageTemplate(usageTemplateWithFlags)
	policyCmd.AddCommand(policyDownloadCmd)
}
