package wafv2

import (
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ManagedRuleGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_wafv2_managed_rule_groups",
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_ManagedRuleGroupSummary.html`,
		Resolver:    fetchWafv2ManagedRuleGroups,
		Transform:   transformers.TransformWithStruct(&types.ManagedRuleGroupSummary{}, transformers.WithPrimaryKeys("Name", "VendorName")),
		Multiplex:   client.ServiceAccountRegionScopeMultiplexer("waf-regional"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "scope",
				Type:     schema.TypeString,
				Resolver: client.ResolveWAFScope,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "properties",
				Type:     schema.TypeJSON,
				Resolver: resolveManageRuleGroupProperties,
			},
		},
	}
}
