package waf

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

type WebACLWrapper struct {
	*types.WebACL
	LoggingConfiguration *types.LoggingConfiguration
}

func WafWebAcls() *schema.Table {
	return &schema.Table{
		Name:         "aws_waf_web_acls",
		Description:  "This is AWS WAF Classic documentation",
		Resolver:     fetchWafWebAcls,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveWafWebACLTags,
			},
			{
				Name:        "default_action_type",
				Description: "Specifies how you want AWS WAF to respond to requests that match the settings in a Rule",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DefaultAction.Type"),
			},
			{
				Name:        "id",
				Description: "A unique identifier for a WebACL",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WebACLId"),
			},
			{
				Name:        "metric_name",
				Description: "A friendly name or description for the metrics for this WebACL",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "A friendly name or description of the WebACL",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "Tha Amazon Resource Name (ARN) of the web ACL.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WebACLArn"),
			},
			{
				Name:        "logging_configuration",
				Description: "The LoggingConfiguration for the specified web ACL.",
				Type:        schema.TypeStringArray,
				Resolver:    resolveWafWebACLRuleLoggingConfiguration,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_waf_web_acl_rules",
				Description: "This is AWS WAF Classic documentation",
				Resolver:    fetchWafWebAclRules,
				Columns: []schema.Column{
					{
						Name:        "web_acl_cq_id",
						Description: "Unique CloudQuery ID of aws_waf_web_acls table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "priority",
						Description: "Specifies the order in which the Rules in a WebACL are evaluated",
						Type:        schema.TypeInt,
					},
					{
						Name:        "rule_id",
						Description: "The RuleId for a Rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "action_type",
						Description: "Specifies how you want AWS WAF to respond to requests that match the settings in a Rule",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Action.Type"),
					},
					{
						Name:        "excluded_rules",
						Description: "An array of rules to exclude from a rule group",
						Type:        schema.TypeStringArray,
						Resolver:    resolveWafWebACLRuleExcludedRules,
					},
					{
						Name:        "override_action_type",
						Description: "COUNT overrides the action specified by the individual rule within a RuleGroup . If set to NONE, the rule's action will take place.  ",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("OverrideAction.Type"),
					},
					{
						Name:        "type",
						Description: "The rule type, either REGULAR, as defined by Rule, RATE_BASED, as defined by RateBasedRule, or GROUP, as defined by RuleGroup",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_waf_web_acl_logging_configuration",
				Description: "The LoggingConfiguration for the specified web ACL.",
				Resolver:    fetchWafWebACLLoggingConfiguration,
				Columns: []schema.Column{
					{
						Name:        "web_acl_cq_id",
						Description: "Unique CloudQuery ID of aws_waf_web_acls table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "log_destination_configs",
						Description: "An array of Amazon Kinesis Data Firehose ARNs.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "resource_arn",
						Description: "The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.",
						Type:        schema.TypeString,
					},
					{
						Name:        "redacted_fields",
						Description: "The parts of the request that you want redacted from the logs. For example, if you redact the cookie field, the cookie field in the firehose will be xxx.",
						Type:        schema.TypeJSON,
						Resolver:    resolveWafWebACLLoggingConfigurationRedactedFields,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchWafWebAcls(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().Waf
	config := waf.ListWebACLsInput{}
	for {
		output, err := service.ListWebACLs(ctx, &config, func(options *waf.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, webAcl := range output.WebACLs {
			webAclConfig := waf.GetWebACLInput{WebACLId: webAcl.WebACLId}
			webAclOutput, err := service.GetWebACL(ctx, &webAclConfig, func(options *waf.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return diag.WrapError(err)
			}

			cfg := waf.GetLoggingConfigurationInput{
				ResourceArn: webAclOutput.WebACL.WebACLArn,
			}
			loggingConfigurationOutput, err := service.GetLoggingConfiguration(ctx, &cfg, func(options *waf.Options) {
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
func resolveWafWebACLTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL := resource.Item.(*WebACLWrapper)

	// Resolve tags for resource
	awsClient := meta.(*client.Client)
	service := awsClient.Services().Waf
	outputTags := make(map[string]*string)
	tagsConfig := waf.ListTagsForResourceInput{ResourceARN: webACL.WebACLArn}
	for {
		tags, err := service.ListTagsForResource(ctx, &tagsConfig, func(options *waf.Options) {
			options.Region = awsClient.Region
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
	return resource.Set("tags", outputTags)
}
func fetchWafWebAclRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	webACL := parent.Item.(*WebACLWrapper)
	res <- webACL.Rules
	return nil
}
func resolveWafWebACLRuleExcludedRules(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule := resource.Item.(types.ActivatedRule)
	excludedRules := make([]string, len(rule.ExcludedRules))
	for i := range rule.ExcludedRules {
		excludedRules[i] = aws.ToString(rule.ExcludedRules[i].RuleId)
	}
	return resource.Set(c.Name, excludedRules)
}
func fetchWafWebACLLoggingConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- interface{}) error {
	rule := resource.Item.(*WebACLWrapper)
	res <- rule.LoggingConfiguration
	return nil
}
func resolveWafWebACLLoggingConfigurationRedactedFields(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	if conf := resource.Item.(*types.LoggingConfiguration); conf != nil {
		out, err := json.Marshal(conf.RedactedFields)
		if err != nil {
			return diag.WrapError(err)
		}
		return resource.Set(c.Name, out)
	}
	return nil
}

func resolveWafWebACLRuleLoggingConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	if rule := resource.Item.(*WebACLWrapper); rule.LoggingConfiguration != nil {
		return resource.Set(c.Name, rule.LoggingConfiguration.LogDestinationConfigs)
	}
	return nil
}
