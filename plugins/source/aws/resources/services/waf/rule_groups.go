package waf

import (
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RuleGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_waf_rule_groups",
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_RuleGroupSummary.html`,
		Resolver:    fetchWafRuleGroups,
		Transform:   transformers.TransformWithStruct(&types.RuleGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("waf"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveWafRuleGroupArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveWafRuleGroupTags,
			},
			{
				Name:     "rule_ids",
				Type:     schema.TypeStringArray,
				Resolver: resolveWafRuleGroupRuleIds,
			},
		},
	}
}
