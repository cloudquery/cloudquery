package recipes

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/wafv2/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WAFv2Resources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "ipsets",
			Struct:              &types.IPSet{},
			Description:         "https://docs.aws.amazon.com/waf/latest/APIReference/API_IPSet.html",
			SkipFields:          []string{"Addresses", "ARN"},
			PreResourceResolver: "getIpset",
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
			SubService:  "managed_rule_groups",
			Struct:      &types.ManagedRuleGroupSummary{},
			Description: "https://docs.aws.amazon.com/waf/latest/APIReference/API_ManagedRuleGroupSummary.html",
			SkipFields:  []string{"Scope"},
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
					Name:     "properties",
					Type:     schema.TypeJSON,
					Resolver: "resolveManageRuleGroupProperties",
				},
			},
		},
		{
			SubService:          "regex_pattern_sets",
			Struct:              &types.RegexPatternSet{},
			Description:         "https://docs.aws.amazon.com/waf/latest/APIReference/API_RegexPatternSet.html",
			SkipFields:          []string{"ARN"},
			PreResourceResolver: "getRegexPatternSet",
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
			SubService:          "rule_groups",
			Struct:              &types.RuleGroup{},
			Description:         "https://docs.aws.amazon.com/waf/latest/APIReference/API_RuleGroup.html",
			SkipFields:          []string{"ARN"},
			PreResourceResolver: "getRuleGroup",
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
			SubService:          "web_acls",
			Struct:              &models.WebACLWrapper{},
			SkipFields:          []string{"ARN"},
			PreResourceResolver: "getWebAcl",
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
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
