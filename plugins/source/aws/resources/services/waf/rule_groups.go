package waf

import (
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RuleGroups() *schema.Table {
	tableName := "aws_waf_rule_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_RuleGroupSummary.html`,
		Resolver:    fetchWafRuleGroups,
		Transform:   client.TransformWithStruct(&types.RuleGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "waf"),
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
