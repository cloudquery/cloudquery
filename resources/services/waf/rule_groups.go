package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func WafRuleGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_waf_rule_groups",
		Description:  "This is AWS WAF Classic documentation",
		Resolver:     fetchWafRuleGroups,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreCommonErrors,
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveWafRuleGroupArn,
			},
			{
				Name:          "rule_ids",
				Type:          schema.TypeStringArray,
				Resolver:      resolveWafRuleGroupRuleIds,
				IgnoreInTests: true,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveWafRuleGroupTags,
			},
			{
				Name:        "id",
				Description: "A unique identifier for a RuleGroup",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RuleGroupId"),
			},
			{
				Name:        "metric_name",
				Description: "A friendly name or description for the metrics for this RuleGroup",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The friendly name or description for the RuleGroup",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchWafRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().Waf
	config := waf.ListRuleGroupsInput{}
	for {
		output, err := service.ListRuleGroups(ctx, &config, func(options *waf.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, rG := range output.RuleGroups {
			ruleGroup, err := service.GetRuleGroup(ctx, &waf.GetRuleGroupInput{RuleGroupId: rG.RuleGroupId}, func(options *waf.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return diag.WrapError(err)
			}
			res <- ruleGroup.RuleGroup
		}

		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.NextMarker = output.NextMarker
	}
	return nil
}
func resolveWafRuleGroupArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	ruleGroup := resource.Item.(*types.RuleGroup)
	return diag.WrapError(resource.Set(c.Name, cl.ARN("waf", "rulegroup", aws.ToString(ruleGroup.RuleGroupId))))
}
func resolveWafRuleGroupRuleIds(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ruleGroup := resource.Item.(*types.RuleGroup)

	// Resolves rule group rules
	awsClient := meta.(*client.Client)
	service := awsClient.Services().Waf
	listActivatedRulesConfig := waf.ListActivatedRulesInRuleGroupInput{RuleGroupId: ruleGroup.RuleGroupId}
	var ruleIDs []string
	for {
		rules, err := service.ListActivatedRulesInRuleGroup(ctx, &listActivatedRulesConfig, func(options *waf.Options) {
			options.Region = awsClient.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, rule := range rules.ActivatedRules {
			ruleIDs = append(ruleIDs, aws.ToString(rule.RuleId))
		}

		if aws.ToString(rules.NextMarker) == "" {
			break
		}
		listActivatedRulesConfig.NextMarker = rules.NextMarker
	}
	return diag.WrapError(resource.Set("rule_ids", ruleIDs))
}
func resolveWafRuleGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ruleGroup := resource.Item.(*types.RuleGroup)

	// Resolve tags for resource
	cl := meta.(*client.Client)
	service := cl.Services().Waf

	// Generate arn
	arnStr := cl.AccountGlobalARN("waf", "rulegroup", aws.ToString(ruleGroup.RuleGroupId))

	outputTags := make(map[string]*string)
	tagsConfig := waf.ListTagsForResourceInput{ResourceARN: aws.String(arnStr)}
	for {
		tags, err := service.ListTagsForResource(ctx, &tagsConfig, func(options *waf.Options) {
			// Set region to default global region
			options.Region = "us-east-1"
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
	return diag.WrapError(resource.Set("tags", outputTags))
}
