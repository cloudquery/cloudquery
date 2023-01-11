package wafv2

import (
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RuleGroups() *schema.Table {
	return &schema.Table{
		Name:                "aws_wafv2_rule_groups",
		Description:         `https://docs.aws.amazon.com/waf/latest/APIReference/API_RuleGroup.html`,
		Resolver:            fetchWafv2RuleGroups,
		PreResourceResolver: getRuleGroup,
		Transform:           transformers.TransformWithStruct(&types.RuleGroup{}),
		Multiplex:           client.ServiceAccountRegionScopeMultiplexer("waf-regional"),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRuleGroupTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "policy",
				Type:     schema.TypeJSON,
				Resolver: resolveWafv2ruleGroupPolicy,
			},
		},
	}
}
