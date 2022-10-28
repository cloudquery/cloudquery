package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WAFRegionalResources() []*Resource {
	resources := []*Resource{

		{
			SubService:  "rate_based_rules",
			Struct:      &types.RateBasedRule{},
			Description: "https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_RateBasedRule.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveWafregionalRateBasedRuleArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveWafregionalRateBasedRuleTags`,
					},
				}...),
		},
		{
			SubService:  "rule_groups",
			Struct:      &types.RuleGroup{},
			Description: "https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_RuleGroup.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveWafregionalRuleGroupArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:        "tags",
						Description: "Rule group tags.",
						Type:        schema.TypeJSON,
						Resolver:    `resolveWafregionalRuleGroupTags`,
					},
				}...),
		},
		{
			SubService:  "rules",
			Struct:      &types.Rule{},
			Description: "https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_Rule.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveWafregionalRuleArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:        "tags",
						Description: "Rule tags.",
						Type:        schema.TypeJSON,
						Resolver:    `resolveWafregionalRuleTags`,
					},
				}...),
		},
		{
			SubService:  "web_acls",
			Struct:      &types.WebACL{},
			Description: "https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_WebACL.html",
			SkipFields:  []string{"WebACLArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("WebACLArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:        "tags",
						Description: "Web ACL tags.",
						Type:        schema.TypeJSON,
						Resolver:    `resolveWafregionalWebACLTags`,
					},
					{
						Name:     "resources_for_web_acl",
						Type:     schema.TypeStringArray,
						Resolver: `resolveWafregionalWebACLResourcesForWebACL`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "wafregional"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("waf-regional")`
	}
	return resources
}
