package waf

import (
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SubscribedRuleGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_waf_subscribed_rule_groups",
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_SubscribedRuleGroupSummary.html`,
		Resolver:    fetchWafSubscribedRuleGroups,
		Transform:   transformers.TransformWithStruct(&types.SubscribedRuleGroupSummary{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("waf"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
				Description: `The AWS Account ID of the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "rule_group_id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RuleGroupId"),
				Description: `A unique identifier for a RuleGroup.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
