package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Wafv2WebAcls() *schema.Table {
	return &schema.Table{
		Name:         "aws_wafv2_web_acls",
		Resolver:     fetchWafv2WebAcls,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name:     "resources_for_web_acl",
				Type:     schema.TypeStringArray,
				Resolver: resolveWafv2webACLResourcesForWebACL,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveWafv2webACLTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
			},
			{
				Name:     "default_action",
				Type:     schema.TypeJSON,
				Resolver: resolveWafv2webACLDefaultAction,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name:     "visibility_config_cloud_watch_metrics_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("VisibilityConfig.CloudWatchMetricsEnabled"),
			},
			{
				Name:     "visibility_config_metric_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VisibilityConfig.MetricName"),
			},
			{
				Name:     "visibility_config_sampled_requests_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("VisibilityConfig.SampledRequestsEnabled"),
			},
			{
				Name: "capacity",
				Type: schema.TypeBigInt,
			},
			{
				Name: "custom_response_bodies",
				Type: schema.TypeJSON,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "label_namespace",
				Type: schema.TypeString,
			},
			{
				Name: "managed_by_firewall_manager",
				Type: schema.TypeBool,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_wafv2_web_acl_rules",
				Resolver: fetchWafv2WebAclRules,
				Columns: []schema.Column{
					{
						Name:     "web_acl_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "priority",
						Type: schema.TypeInt,
					},
					{
						Name:     "statement",
						Type:     schema.TypeJSON,
						Resolver: resolveWafv2webACLRuleStatement,
					},
					{
						Name:     "visibility_config_cloud_watch_metrics_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VisibilityConfig.CloudWatchMetricsEnabled"),
					},
					{
						Name:     "visibility_config_metric_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VisibilityConfig.MetricName"),
					},
					{
						Name:     "visibility_config_sampled_requests_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VisibilityConfig.SampledRequestsEnabled"),
					},
					{
						Name:     "action",
						Type:     schema.TypeJSON,
						Resolver: resolveWafv2webACLRuleAction,
					},
					{
						Name:     "override_action",
						Type:     schema.TypeJSON,
						Resolver: resolveWafv2webACLRuleOverrideAction,
					},
					{
						Name:     "labels",
						Type:     schema.TypeStringArray,
						Resolver: resolveWafv2webACLRuleLabels,
					},
				},
			},
			{
				Name:     "aws_wafv2_web_acl_post_process_firewall_manager_rule_groups",
				Resolver: fetchWafv2WebAclPostProcessFirewallManagerRuleGroups,
				Columns: []schema.Column{
					{
						Name:     "web_acl_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "statement",
						Type:     schema.TypeJSON,
						Resolver: resolveWafv2webACLPostProcessFirewallManagerRuleGroupStatement,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name:     "override_action",
						Type:     schema.TypeJSON,
						Resolver: resolveWafv2webACLPostProcessFirewallManagerRuleGroupOverrideAction,
					},
					{
						Name: "priority",
						Type: schema.TypeInt,
					},
					{
						Name:     "visibility_config_cloud_watch_metrics_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VisibilityConfig.CloudWatchMetricsEnabled"),
					},
					{
						Name:     "visibility_config_metric_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VisibilityConfig.MetricName"),
					},
					{
						Name:     "visibility_config_sampled_requests_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VisibilityConfig.SampledRequestsEnabled"),
					},
				},
			},
			{
				Name:     "aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups",
				Resolver: fetchWafv2WebAclPreProcessFirewallManagerRuleGroups,
				Columns: []schema.Column{
					{
						Name:     "web_acl_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "statement",
						Type:     schema.TypeJSON,
						Resolver: resolveWafv2webACLPreProcessFirewallManagerRuleGroupStatement,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name:     "override_action",
						Type:     schema.TypeJSON,
						Resolver: resolveWafv2webACLPreProcessFirewallManagerRuleGroupOverrideAction,
					},
					{
						Name: "priority",
						Type: schema.TypeInt,
					},
					{
						Name:     "visibility_config_cloud_watch_metrics_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VisibilityConfig.CloudWatchMetricsEnabled"),
					},
					{
						Name:     "visibility_config_metric_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VisibilityConfig.MetricName"),
					},
					{
						Name:     "visibility_config_sampled_requests_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VisibilityConfig.SampledRequestsEnabled"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchWafv2WebAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().WafV2
	config := wafv2.ListWebACLsInput{}
	for {
		output, err := service.ListWebACLs(ctx, &config, func(options *wafv2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, webAcl := range output.WebACLs {
			webAclConfig := wafv2.GetWebACLInput{Id: webAcl.Id}
			webAclOutput, err := service.GetWebACL(ctx, &webAclConfig, func(options *wafv2.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- webAclOutput.WebACL
		}

		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.NextMarker = output.NextMarker
	}
	return nil
}

func resolveWafv2webACLResourcesForWebACL(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL, ok := resource.Item.(*types.WebACL)
	if !ok {
		return fmt.Errorf("not an WebACL instance: %#v", resource.Item)
	}

	client := meta.(*client.Client)
	service := client.Services().WafV2

	// Resolve resources that are associated with the given web ACL
	resourceArns, err := service.ListResourcesForWebACL(ctx, &wafv2.ListResourcesForWebACLInput{WebACLArn: webACL.ARN}, func(options *wafv2.Options) {
		options.Region = client.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, resourceArns.ResourceArns)
}

func resolveWafv2webACLTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL, ok := resource.Item.(*types.WebACL)
	if !ok {
		return fmt.Errorf("not an WebACL instance: %#v", resource.Item)
	}

	client := meta.(*client.Client)
	service := client.Services().WafV2

	// Resolve tags
	outputTags := make(map[string]*string)
	tagsConfig := wafv2.ListTagsForResourceInput{ResourceARN: webACL.ARN}
	for {
		tags, err := service.ListTagsForResource(ctx, &tagsConfig, func(options *wafv2.Options) {
			options.Region = client.Region
		})
		if err != nil {
			return err
		}
		for _, t := range tags.TagInfoForResource.TagList {
			outputTags[*t.Key] = t.Value
		}
		if aws.ToString(tags.NextMarker) == "" {
			break
		}
		tagsConfig.NextMarker = tags.NextMarker
	}
	return resource.Set(c.Name, outputTags)
}

func resolveWafv2webACLDefaultAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL, ok := resource.Item.(*types.WebACL)
	if !ok {
		return fmt.Errorf("not an WebACL instance: %#v", resource.Item)
	}
	if webACL.DefaultAction == nil {
		return nil
	}
	data, err := json.Marshal(webACL.DefaultAction)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func fetchWafv2WebAclRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	webACL, ok := parent.Item.(*types.WebACL)
	if !ok {
		return fmt.Errorf("not an WebACL instance: %#v", parent.Item)
	}
	res <- webACL.Rules
	return nil
}

func resolveWafv2webACLRuleStatement(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule, ok := resource.Item.(types.Rule)
	if !ok {
		return fmt.Errorf("not an Rule instance: %#v", resource.Item)
	}
	if rule.Statement == nil {
		return nil
	}
	data, err := json.Marshal(rule.Statement)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func resolveWafv2webACLRuleAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule, ok := resource.Item.(types.Rule)
	if !ok {
		return fmt.Errorf("not an Rule instance: %#v", resource.Item)
	}
	if rule.Action == nil {
		return nil
	}
	data, err := json.Marshal(rule.Action)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func resolveWafv2webACLRuleOverrideAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule, ok := resource.Item.(types.Rule)
	if !ok {
		return fmt.Errorf("not an Rule instance: %#v", resource.Item)
	}
	if rule.OverrideAction == nil {
		return nil
	}
	data, err := json.Marshal(rule.OverrideAction)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func resolveWafv2webACLRuleLabels(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule, ok := resource.Item.(types.Rule)
	if !ok {
		return fmt.Errorf("not an Rule instance: %#v", resource.Item)
	}
	labels := make([]string, len(rule.RuleLabels))
	for i := range rule.RuleLabels {
		labels[i] = aws.ToString(rule.RuleLabels[i].Name)
	}
	return resource.Set(c.Name, labels)
}

func fetchWafv2WebAclPostProcessFirewallManagerRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	webACL, ok := parent.Item.(*types.WebACL)
	if !ok {
		return fmt.Errorf("not an WebACL instance: %#v", parent.Item)
	}
	res <- webACL.PostProcessFirewallManagerRuleGroups
	return nil
}

func resolveWafv2webACLPostProcessFirewallManagerRuleGroupStatement(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	firewallManagerRuleGroup, ok := resource.Item.(types.FirewallManagerRuleGroup)
	if !ok {
		return fmt.Errorf("not an FirewallManagerRuleGroup instance: %#v", resource.Item)
	}
	if firewallManagerRuleGroup.FirewallManagerStatement == nil {
		return nil
	}
	data, err := json.Marshal(firewallManagerRuleGroup.FirewallManagerStatement)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func resolveWafv2webACLPostProcessFirewallManagerRuleGroupOverrideAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	firewallManagerRuleGroup, ok := resource.Item.(types.FirewallManagerRuleGroup)
	if !ok {
		return fmt.Errorf("not an FirewallManagerRuleGroup instance: %#v", resource.Item)
	}
	if firewallManagerRuleGroup.OverrideAction == nil {
		return nil
	}
	data, err := json.Marshal(firewallManagerRuleGroup.OverrideAction)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func fetchWafv2WebAclPreProcessFirewallManagerRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	webACL, ok := parent.Item.(*types.WebACL)
	if !ok {
		return fmt.Errorf("not an WebACL instance: %#v", parent.Item)
	}
	res <- webACL.PreProcessFirewallManagerRuleGroups
	return nil
}

func resolveWafv2webACLPreProcessFirewallManagerRuleGroupStatement(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	firewallManagerRuleGroup, ok := resource.Item.(types.FirewallManagerRuleGroup)
	if !ok {
		return fmt.Errorf("not an FirewallManagerRuleGroup instance: %#v", resource.Item)
	}
	if firewallManagerRuleGroup.FirewallManagerStatement == nil {
		return nil
	}
	data, err := json.Marshal(firewallManagerRuleGroup.FirewallManagerStatement)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func resolveWafv2webACLPreProcessFirewallManagerRuleGroupOverrideAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	firewallManagerRuleGroup, ok := resource.Item.(types.FirewallManagerRuleGroup)
	if !ok {
		return fmt.Errorf("not an FirewallManagerRuleGroup instance: %#v", resource.Item)
	}
	if firewallManagerRuleGroup.OverrideAction == nil {
		return nil
	}
	data, err := json.Marshal(firewallManagerRuleGroup.OverrideAction)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
