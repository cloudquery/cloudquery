// Code generated by codegen; DO NOT EDIT.

package wafv2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WebAcls() *schema.Table {
	return &schema.Table{
		Name:      "aws_wafv2_web_acls",
		Resolver:  fetchWafv2WebAcls,
		Multiplex: client.ServiceAccountRegionScopeMultiplexer("waf-regional"),
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
				Resolver: resolveWebACLTags,
			},
			{
				Name:     "resources_for_web_acl",
				Type:     schema.TypeStringArray,
				Resolver: resolveWafv2webACLResourcesForWebACL,
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
				Name:     "default_action",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultAction"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "visibility_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VisibilityConfig"),
			},
			{
				Name:     "capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Capacity"),
			},
			{
				Name:     "captcha_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CaptchaConfig"),
			},
			{
				Name:     "custom_response_bodies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CustomResponseBodies"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "label_namespace",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LabelNamespace"),
			},
			{
				Name:     "managed_by_firewall_manager",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ManagedByFirewallManager"),
			},
			{
				Name:     "post_process_firewall_manager_rule_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PostProcessFirewallManagerRuleGroups"),
			},
			{
				Name:     "pre_process_firewall_manager_rule_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PreProcessFirewallManagerRuleGroups"),
			},
			{
				Name:     "rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Rules"),
			},
		},
	}
}
