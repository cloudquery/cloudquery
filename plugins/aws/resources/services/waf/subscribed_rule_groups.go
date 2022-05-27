package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func WafSubscribedRuleGroups() *schema.Table {
	return &schema.Table{
		Name:          "aws_waf_subscribed_rule_groups",
		Description:   "This is AWS WAF Classic documentation",
		Resolver:      fetchWafSubscribedRuleGroups,
		Multiplex:     client.AccountMultiplex,
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountFilter,
		IgnoreInTests: true,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "rule_group_id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "metric_name",
				Description: "A friendly name or description for the metrics for this RuleGroup",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "A friendly name or description of the RuleGroup",
				Type:        schema.TypeString,
			},
			{
				Name:        "rule_group_id",
				Description: "A unique identifier for a RuleGroup.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchWafSubscribedRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().Waf
	config := waf.ListSubscribedRuleGroupsInput{}
	for {
		output, err := service.ListSubscribedRuleGroups(ctx, &config, func(options *waf.Options) {
			// Set region to default global region
			options.Region = "us-east-1"
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.RuleGroups

		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.NextMarker = output.NextMarker
	}
	return nil
}
