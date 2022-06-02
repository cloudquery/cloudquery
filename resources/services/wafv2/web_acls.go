package wafv2

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

type WebACLWrapper struct {
	*types.WebACL
	LoggingConfiguration *types.LoggingConfiguration
}

func Wafv2WebAcls() *schema.Table {
	return &schema.Table{
		Name:         "aws_wafv2_web_acls",
		Description:  "A Web ACL defines a collection of rules to use to inspect and control web requests",
		Resolver:     fetchWafv2WebAcls,
		Multiplex:    client.ServiceAccountRegionScopeMultiplexer("waf-regional"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionScopeFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
				Name:        "scope",
				Description: "Specifies whether this is for an Amazon CloudFront distribution or for a regional application.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveWAFScope,
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.  ",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ARN"),
			},
			{
				Name:        "default_action",
				Description: "The action to perform if none of the Rules contained in the WebACL match.  ",
				Type:        schema.TypeJSON,
				Resolver:    resolveWafv2webACLDefaultAction,
			},
			{
				Name:        "id",
				Description: "A unique identifier for the WebACL",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
			},
			{
				Name:        "name",
				Description: "The name of the Web ACL",
				Type:        schema.TypeString,
			},
			{
				Name:        "visibility_config_cloud_watch_metrics_enabled",
				Description: "A boolean indicating whether the associated resource sends metrics to CloudWatch",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VisibilityConfig.CloudWatchMetricsEnabled"),
			},
			{
				Name:        "visibility_config_metric_name",
				Description: "A name of the CloudWatch metric",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VisibilityConfig.MetricName"),
			},
			{
				Name:        "visibility_config_sampled_requests_enabled",
				Description: "A boolean indicating whether AWS WAF should store a sampling of the web requests that match the rules",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VisibilityConfig.SampledRequestsEnabled"),
			},
			{
				Name:        "capacity",
				Description: "The web ACL capacity units (WCUs) currently being used by this web ACL",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "custom_response_bodies",
				Description: "A map of custom response keys and content bodies",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "description",
				Description: "A description of the Web ACL that helps with identification.",
				Type:        schema.TypeString,
			},
			{
				Name:        "label_namespace",
				Description: "The label namespace prefix for this web ACL",
				Type:        schema.TypeString,
			},
			{
				Name:        "managed_by_firewall_manager",
				Description: "Indicates whether this web ACL is managed by AWS Firewall Manager",
				Type:        schema.TypeBool,
			},
			{
				Name:        "logging_configuration",
				Description: "The LoggingConfiguration for the specified web ACL.",
				Type:        schema.TypeStringArray,
				Resolver:    resolveWafV2WebACLRuleLoggingConfiguration,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_wafv2_web_acl_rules",
				Description: "A single rule, which you can use in a WebACL or RuleGroup to identify web requests that you want to allow, block, or count",
				Resolver:    fetchWafv2WebAclRules,
				Columns: []schema.Column{
					{
						Name:        "web_acl_cq_id",
						Description: "Unique CloudQuery ID of aws_wafv2_web_acls table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The name of the rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "priority",
						Description: "If you define more than one Rule in a WebACL, AWS WAF evaluates each request against the Rules in order based on the value of Priority",
						Type:        schema.TypeInt,
					},
					{
						Name:        "statement",
						Description: "The AWS WAF processing statement for the rule, for example ByteMatchStatement or SizeConstraintStatement.  ",
						Type:        schema.TypeJSON,
						Resolver:    resolveWafv2webACLRuleStatement,
					},
					{
						Name:        "visibility_config_cloud_watch_metrics_enabled",
						Description: "A boolean indicating whether the associated resource sends metrics to CloudWatch",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VisibilityConfig.CloudWatchMetricsEnabled"),
					},
					{
						Name:        "visibility_config_metric_name",
						Description: "A name of the CloudWatch metric",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VisibilityConfig.MetricName"),
					},
					{
						Name:        "visibility_config_sampled_requests_enabled",
						Description: "A boolean indicating whether AWS WAF should store a sampling of the web requests that match the rules",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VisibilityConfig.SampledRequestsEnabled"),
					},
					{
						Name:        "action",
						Description: "The action that AWS WAF should take on a web request when it matches the rule statement",
						Type:        schema.TypeJSON,
						Resolver:    resolveWafv2webACLRuleAction,
					},
					{
						Name:        "override_action",
						Description: "The override action to apply to the rules in a rule group",
						Type:        schema.TypeJSON,
						Resolver:    resolveWafv2webACLRuleOverrideAction,
					},
					{
						Name:        "labels",
						Description: "Labels to apply to web requests that match the rule match statement",
						Type:        schema.TypeStringArray,
						Resolver:    resolveWafv2webACLRuleLabels,
					},
				},
			},
			{
				Name:          "aws_wafv2_web_acl_post_process_firewall_manager_rule_groups",
				Description:   "A rule group that's defined for an AWS Firewall Manager WAF policy. ",
				Resolver:      fetchWafv2WebAclPostProcessFirewallManagerRuleGroups,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "web_acl_cq_id",
						Description: "Unique CloudQuery ID of aws_wafv2_web_acls table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "statement",
						Description: "The processing guidance for an AWS Firewall Manager rule",
						Type:        schema.TypeJSON,
						Resolver:    resolveWafv2webACLPostProcessFirewallManagerRuleGroupStatement,
					},
					{
						Name:        "name",
						Description: "The name of the rule group",
						Type:        schema.TypeString,
					},
					{
						Name:        "override_action",
						Description: "The override action to apply to the rules in a rule group",
						Type:        schema.TypeJSON,
						Resolver:    resolveWafv2webACLPostProcessFirewallManagerRuleGroupOverrideAction,
					},
					{
						Name:        "priority",
						Description: "If you define more than one rule group in the first or last Firewall Manager rule groups, AWS WAF evaluates each request against the rule groups in order, starting from the lowest priority setting",
						Type:        schema.TypeInt,
					},
					{
						Name:        "visibility_config_cloud_watch_metrics_enabled",
						Description: "A boolean indicating whether the associated resource sends metrics to CloudWatch",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VisibilityConfig.CloudWatchMetricsEnabled"),
					},
					{
						Name:        "visibility_config_metric_name",
						Description: "A name of the CloudWatch metric",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VisibilityConfig.MetricName"),
					},
					{
						Name:        "visibility_config_sampled_requests_enabled",
						Description: "A boolean indicating whether AWS WAF should store a sampling of the web requests that match the rules",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VisibilityConfig.SampledRequestsEnabled"),
					},
				},
			},
			{
				Name:          "aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups",
				Description:   "A rule group that's defined for an AWS Firewall Manager WAF policy. ",
				Resolver:      fetchWafv2WebAclPreProcessFirewallManagerRuleGroups,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "web_acl_cq_id",
						Description: "Unique CloudQuery ID of aws_wafv2_web_acls table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "statement",
						Description: "The processing guidance for an AWS Firewall Manager rule",
						Type:        schema.TypeJSON,
						Resolver:    resolveWafv2webACLPreProcessFirewallManagerRuleGroupStatement,
					},
					{
						Name:        "name",
						Description: "The name of the rule group",
						Type:        schema.TypeString,
					},
					{
						Name:        "override_action",
						Description: "The override action to apply to the rules in a rule group",
						Type:        schema.TypeJSON,
						Resolver:    resolveWafv2webACLPreProcessFirewallManagerRuleGroupOverrideAction,
					},
					{
						Name:        "priority",
						Description: "If you define more than one rule group in the first or last Firewall Manager rule groups, AWS WAF evaluates each request against the rule groups in order, starting from the lowest priority setting",
						Type:        schema.TypeInt,
					},
					{
						Name:        "visibility_config_cloud_watch_metrics_enabled",
						Description: "A boolean indicating whether the associated resource sends metrics to CloudWatch",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VisibilityConfig.CloudWatchMetricsEnabled"),
					},
					{
						Name:        "visibility_config_metric_name",
						Description: "A name of the CloudWatch metric",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VisibilityConfig.MetricName"),
					},
					{
						Name:        "visibility_config_sampled_requests_enabled",
						Description: "A boolean indicating whether AWS WAF should store a sampling of the web requests that match the rules",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VisibilityConfig.SampledRequestsEnabled"),
					},
				},
			},
			{
				Name:        "aws_wafv2_web_acl_logging_configuration",
				Description: "The LoggingConfiguration for the specified web ACL.",
				Resolver:    fetchWafv2WebACLLoggingConfiguration,
				Columns: []schema.Column{
					{
						Name:        "web_acl_cq_id",
						Description: "Unique CloudQuery ID of aws_wafv2_web_acls table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "log_destination_configs",
						Description: "The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "resource_arn",
						Description: "The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.",
						Type:        schema.TypeString,
					},
					{
						Name:        "logging_filter",
						Description: "Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.",
						Type:        schema.TypeJSON,
						Resolver:    resolveWafv2WebACLLoggingConfigurationLoggingFilter,
					},
					{
						Name:        "managed_by_firewall_manager",
						Description: "Indicates whether the logging configuration was created by Firewall Manager, as part of an WAF policy configuration. If true, only Firewall Manager can modify or delete the configuration.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "redacted_fields",
						Description: "The parts of the request that you want redacted from the logs. For example, if you redact the cookie field, the cookie field in the firehose will be xxx. You can specify only the following fields for redaction: UriPath, QueryString, SingleHeader, Method, and JsonBody.",
						Type:        schema.TypeJSON,
						Resolver:    resolveWafv2WebACLLoggingConfigurationRedactedFields,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchWafv2WebAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().WafV2

	config := wafv2.ListWebACLsInput{
		Scope: c.WAFScope,
		Limit: aws.Int32(100),
	}
	for {
		output, err := service.ListWebACLs(ctx, &config, func(options *wafv2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, webAcl := range output.WebACLs {
			webAclConfig := wafv2.GetWebACLInput{Id: webAcl.Id, Name: webAcl.Name, Scope: c.WAFScope}
			webAclOutput, err := service.GetWebACL(ctx, &webAclConfig, func(options *wafv2.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return diag.WrapError(err)
			}

			cfg := wafv2.GetLoggingConfigurationInput{
				ResourceArn: webAclOutput.WebACL.ARN,
			}

			loggingConfigurationOutput, err := service.GetLoggingConfiguration(ctx, &cfg, func(options *wafv2.Options) {
				options.Region = c.Region
			})
			if err != nil {
				if client.IsAWSError(err, "WAFNonexistentItemException") {
					c.Logger().Debug("Logging configuration not found for: %s", webAclOutput.WebACL.Name)
				} else {
					c.Logger().Error("GetLoggingConfiguration failed with error: %s", err.Error())
				}
			}

			var webAclLoggingConfiguration *types.LoggingConfiguration
			if loggingConfigurationOutput != nil {
				webAclLoggingConfiguration = loggingConfigurationOutput.LoggingConfiguration
			}

			res <- &WebACLWrapper{
				webAclOutput.WebACL,
				webAclLoggingConfiguration,
			}
		}

		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.NextMarker = output.NextMarker
	}
	return nil
}
func resolveWafv2webACLResourcesForWebACL(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL := resource.Item.(*WebACLWrapper)

	cl := meta.(*client.Client)
	service := cl.Services().WafV2

	resourceArns := []string{}
	if cl.WAFScope == types.ScopeCloudfront {
		cloudfrontService := cl.Services().Cloudfront
		params := &cloudfront.ListDistributionsByWebACLIdInput{
			WebACLId: webACL.Id,
			MaxItems: aws.Int32(100),
		}
		for {
			output, err := cloudfrontService.ListDistributionsByWebACLId(ctx, params, func(options *cloudfront.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return diag.WrapError(err)
			}
			for _, item := range output.DistributionList.Items {
				resourceArns = append(resourceArns, *item.ARN)
			}
			if aws.ToString(output.DistributionList.NextMarker) == "" {
				break
			}
			params.Marker = output.DistributionList.NextMarker
		}
	} else {
		output, err := service.ListResourcesForWebACL(ctx, &wafv2.ListResourcesForWebACLInput{WebACLArn: webACL.ARN}, func(options *wafv2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		resourceArns = output.ResourceArns
	}
	return diag.WrapError(resource.Set(c.Name, resourceArns))
}
func resolveWafv2webACLTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL := resource.Item.(*WebACLWrapper)

	cl := meta.(*client.Client)
	service := cl.Services().WafV2

	// Resolve tags
	outputTags := make(map[string]*string)
	tagsConfig := wafv2.ListTagsForResourceInput{ResourceARN: webACL.ARN}
	for {
		tags, err := service.ListTagsForResource(ctx, &tagsConfig, func(options *wafv2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, t := range tags.TagInfoForResource.TagList {
			outputTags[*t.Key] = t.Value
		}
		if aws.ToString(tags.NextMarker) == "" {
			break
		}
		tagsConfig.NextMarker = tags.NextMarker
	}
	return diag.WrapError(resource.Set(c.Name, outputTags))
}
func resolveWafv2webACLDefaultAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL := resource.Item.(*WebACLWrapper)
	if webACL.DefaultAction == nil {
		return nil
	}
	data, err := json.Marshal(webACL.DefaultAction)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, data))
}
func fetchWafv2WebAclRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	webACL := parent.Item.(*WebACLWrapper)
	res <- webACL.Rules
	return nil
}
func resolveWafv2webACLRuleStatement(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule := resource.Item.(types.Rule)
	if rule.Statement == nil {
		return nil
	}
	data, err := json.Marshal(rule.Statement)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveWafv2webACLRuleAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule := resource.Item.(types.Rule)
	if rule.Action == nil {
		return nil
	}
	data, err := json.Marshal(rule.Action)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveWafv2webACLRuleOverrideAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule := resource.Item.(types.Rule)
	if rule.OverrideAction == nil {
		return nil
	}
	data, err := json.Marshal(rule.OverrideAction)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveWafv2webACLRuleLabels(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule := resource.Item.(types.Rule)
	labels := make([]string, len(rule.RuleLabels))
	for i := range rule.RuleLabels {
		labels[i] = aws.ToString(rule.RuleLabels[i].Name)
	}
	return diag.WrapError(resource.Set(c.Name, labels))
}
func fetchWafv2WebAclPostProcessFirewallManagerRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	webACL := parent.Item.(*WebACLWrapper)
	res <- webACL.PostProcessFirewallManagerRuleGroups
	return nil
}
func resolveWafv2webACLPostProcessFirewallManagerRuleGroupStatement(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	firewallManagerRuleGroup := resource.Item.(types.FirewallManagerRuleGroup)
	if firewallManagerRuleGroup.FirewallManagerStatement == nil {
		return nil
	}
	data, err := json.Marshal(firewallManagerRuleGroup.FirewallManagerStatement)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveWafv2webACLPostProcessFirewallManagerRuleGroupOverrideAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	firewallManagerRuleGroup := resource.Item.(types.FirewallManagerRuleGroup)
	if firewallManagerRuleGroup.OverrideAction == nil {
		return nil
	}
	data, err := json.Marshal(firewallManagerRuleGroup.OverrideAction)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, data))
}
func fetchWafv2WebAclPreProcessFirewallManagerRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	webACL := parent.Item.(*WebACLWrapper)
	res <- webACL.PreProcessFirewallManagerRuleGroups
	return nil
}
func resolveWafv2webACLPreProcessFirewallManagerRuleGroupStatement(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	firewallManagerRuleGroup := resource.Item.(types.FirewallManagerRuleGroup)
	if firewallManagerRuleGroup.FirewallManagerStatement == nil {
		return nil
	}
	data, err := json.Marshal(firewallManagerRuleGroup.FirewallManagerStatement)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, data))
}
func resolveWafv2webACLPreProcessFirewallManagerRuleGroupOverrideAction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	firewallManagerRuleGroup := resource.Item.(types.FirewallManagerRuleGroup)
	if firewallManagerRuleGroup.OverrideAction == nil {
		return nil
	}
	data, err := json.Marshal(firewallManagerRuleGroup.OverrideAction)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, data))
}
func fetchWafv2WebACLLoggingConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- interface{}) error {
	rule := resource.Item.(*WebACLWrapper)
	res <- rule.LoggingConfiguration
	return nil
}
func resolveWafv2WebACLLoggingConfigurationRedactedFields(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	if conf := resource.Item.(*types.LoggingConfiguration); conf != nil {
		out, err := json.Marshal(conf.RedactedFields)
		if err != nil {
			return diag.WrapError(err)
		}
		return diag.WrapError(resource.Set(c.Name, out))
	}
	return nil
}
func resolveWafv2WebACLLoggingConfigurationLoggingFilter(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	if conf := resource.Item.(*types.LoggingConfiguration); conf != nil {
		out, err := json.Marshal(conf.LoggingFilter)
		if err != nil {
			return diag.WrapError(err)
		}
		return diag.WrapError(resource.Set(c.Name, out))
	}
	return nil
}
func resolveWafV2WebACLRuleLoggingConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	if rule := resource.Item.(*WebACLWrapper); rule.LoggingConfiguration != nil {
		return diag.WrapError(resource.Set(c.Name, rule.LoggingConfiguration.LogDestinationConfigs))
	}
	return nil
}
