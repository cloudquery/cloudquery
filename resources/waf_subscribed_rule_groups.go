package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func WafSubscribedRuleGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_waf_subscribed_rule_groups",
		Resolver:     fetchWafSubscribedRuleGroups,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "metric_name",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "rule_group_id",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchWafSubscribedRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().Waf
	config := waf.ListSubscribedRuleGroupsInput{}
	for {
		output, err := service.ListSubscribedRuleGroups(ctx, &config, func(options *waf.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.RuleGroups

		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.NextMarker = output.NextMarker
	}
	return nil
}
