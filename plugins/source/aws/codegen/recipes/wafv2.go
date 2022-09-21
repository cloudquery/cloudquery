package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WAFv2Resources() []*Resource {
	resources := []*Resource{
		{
			SubService: "ipsets",
			Struct:     &types.IPSet{},
			SkipFields: []string{"Addresses", "ARN"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSAccount",
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSRegion",
				},
				{
					Name:     "addresses",
					Type:     schema.TypeInetArray,
					Resolver: "resolveIpsetAddresses",
				},
				{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: "resolveIpsetTags",
				},
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("ARN")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService:           "managed_rule_groups",
			Struct:               &types.ManagedRuleGroupSummary{},
			PostResourceResolver: "resolveDescribeManagedRuleGroup",
			SkipFields:           []string{"Scope"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSAccount",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSRegion",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "scope",
					Type:     schema.TypeString,
					Resolver: "client.ResolveWAFScope",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name: "available_labels",
					Type: schema.TypeStringArray,
				},
				{
					Name: "consumed_labels",
					Type: schema.TypeStringArray,
				},
				{
					Name: "capacity",
					Type: schema.TypeInt,
				},
				{
					Name: "label_namespace",
					Type: schema.TypeString,
				},
				{
					Name: "rules",
					Type: schema.TypeJSON,
				},
			},
		},
		{
			SubService: "regex_pattern_sets",
			Struct:     &types.RegexPatternSet{},
			SkipFields: []string{"ARN"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSAccount",
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSRegion",
				},
				{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: "resolveRegexPatternSetTags",
				},
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("ARN")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService: "rule_groups",
			Struct:     &types.RuleGroup{},
			SkipFields: []string{"ARN"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSAccount",
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSRegion",
				},
				{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: "resolveRuleGroupTags",
				},
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("ARN")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "policy",
					Type:     schema.TypeJSON,
					Resolver: `resolveWafv2ruleGroupPolicy`,
				},
			},
		},
		{
			SubService: "web_acls",
			Struct:     &types.WebACL{},
			SkipFields: []string{"ARN"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSAccount",
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSRegion",
				},
				{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: "resolveWebACLTags",
				},
				{
					Name:     "resources_for_web_acl",
					Type:     schema.TypeStringArray,
					Resolver: `resolveWafv2webACLResourcesForWebACL`,
				},
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("ARN")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "wafv2"
		r.Multiplex = `client.ServiceAccountRegionScopeMultiplexer("waf-regional")`
	}
	return resources
}
