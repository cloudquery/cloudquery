package waf

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type WebACLWrapper struct {
	*types.WebACL
	LoggingConfiguration *types.LoggingConfiguration
}

func WafWebAcls() *schema.Table {
	return &schema.Table{
		Name:        "aws_waf_web_acls",
		Description: "This is AWS WAF Classic documentation",
		Resolver:    fetchWafWebAcls,
		Multiplex:   client.AccountMultiplex,
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
				Name:            "arn",
				Description:     "Tha Amazon Resource Name (ARN) of the web ACL.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("WebACLArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "logging_configuration",
				Description: "The LoggingConfiguration for the specified web ACL.",
				Type:        schema.TypeStringArray,
				Resolver:    resolveWafWebACLRuleLoggingConfiguration,
			},
			{
				Name:        "rules",
				Description: "This is AWS WAF Classic documentation",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "configuration",
				Description: "The LoggingConfiguration for the specified web ACL",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchWafWebAcls(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().Waf
	config := waf.ListWebACLsInput{}
	for {
		output, err := service.ListWebACLs(ctx, &config)
		if err != nil {
			return err
		}
		for _, webAcl := range output.WebACLs {
			webAclConfig := waf.GetWebACLInput{WebACLId: webAcl.WebACLId}
			webAclOutput, err := service.GetWebACL(ctx, &webAclConfig, func(options *waf.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}

			cfg := waf.GetLoggingConfigurationInput{
				ResourceArn: webAclOutput.WebACL.WebACLArn,
			}
			loggingConfigurationOutput, err := service.GetLoggingConfiguration(ctx, &cfg, func(options *waf.Options) {
				options.Region = c.Region
			})
			if err != nil {
				c.Logger().Error().Err(err).Msg("GetLoggingConfiguration failed")
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
		tags, err := service.ListTagsForResource(ctx, &tagsConfig)
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
	return resource.Set("tags", outputTags)
}

func resolveWafWebACLRuleExcludedRules(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule := resource.Item.(types.ActivatedRule)
	excludedRules := make([]string, len(rule.ExcludedRules))
	for i := range rule.ExcludedRules {
		excludedRules[i] = aws.ToString(rule.ExcludedRules[i].RuleId)
	}
	return resource.Set(c.Name, excludedRules)
}

func resolveWafWebACLLoggingConfigurationRedactedFields(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	if conf := resource.Item.(*types.LoggingConfiguration); conf != nil {
		out, err := json.Marshal(conf.RedactedFields)
		if err != nil {
			return err
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
