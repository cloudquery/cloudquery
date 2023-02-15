package wafregional

import (
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RuleGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_wafregional_rule_groups",
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_RuleGroup.html`,
		Resolver:    fetchWafregionalRuleGroups,
		Transform:   transformers.TransformWithStruct(&types.RuleGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("waf-regional"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveWafregionalRuleGroupArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveWafregionalRuleGroupTags,
				Description: `Rule group tags.`,
			},
		},
	}
}
