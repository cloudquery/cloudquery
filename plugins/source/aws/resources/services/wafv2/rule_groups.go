package wafv2

import (
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RuleGroups() *schema.Table {
	tableName := "aws_wafv2_rule_groups"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/waf/latest/APIReference/API_RuleGroup.html`,
		Resolver:            fetchWafv2RuleGroups,
		PreResourceResolver: getRuleGroup,
		Transform:           client.TransformWithStruct(&types.RuleGroup{}),
		Multiplex:           client.ServiceAccountRegionScopeMultiplexer(tableName, "waf-regional"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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
