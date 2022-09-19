package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WAFResources() []*Resource {
	resources := []*Resource{

		{
			SubService: "rule_groups",
			Struct:     &types.RuleGroupSummary{},
			SkipFields: []string{"ARN"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveWafRuleGroupArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveWafRuleGroupTags`,
					},
					{
						Name:     "rule_ids",
						Type:     schema.TypeStringArray,
						Resolver: `resolveWafRuleGroupRuleIds`,
					},
				}...),
		},

		{
			SubService: "rules",
			Struct:     &types.RuleSummary{},
			SkipFields: []string{"ARN"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveWafRuleArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveWafRuleTags`,
					},
				}...),
		},

		{
			SubService: "subscribed_rule_groups",
			Struct:     &types.SubscribedRuleGroupSummary{},
			SkipFields: []string{"RuleGroupId"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:        "account_id",
					Description: "The AWS Account ID of the resource.",
					Type:        schema.TypeString,
					Resolver:    `client.ResolveAWSAccount`,
					Options:     schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:        "rule_group_id",
					Description: "A unique identifier for a RuleGroup.",
					Type:        schema.TypeString,
					Options:     schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},

		{
			SubService: "web_acls",
			Struct:     &types.WebACLSummary{},
			SkipFields: []string{"WebACLArn"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("WebACLArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveWafWebACLTags`,
					},
					{
						Name:        "logging_configuration",
						Description: "The LoggingConfiguration for the specified web ACL.",
						Type:        schema.TypeJSON,
						Resolver:    `schema.PathResolver("LoggingConfiguration")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "waf"
		r.Multiplex = `client.AccountMultiplex`
	}
	return resources
}
