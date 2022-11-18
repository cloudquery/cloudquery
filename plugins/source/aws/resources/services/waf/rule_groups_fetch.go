package waf

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchWafRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().Waf
	config := waf.ListRuleGroupsInput{}
	for {
		output, err := service.ListRuleGroups(ctx, &config)
		if err != nil {
			return err
		}
		for _, rG := range output.RuleGroups {
			ruleGroup, err := service.GetRuleGroup(ctx, &waf.GetRuleGroupInput{RuleGroupId: rG.RuleGroupId}, func(options *waf.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
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
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   "waf",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("rulegroup/%s", aws.ToString(ruleGroup.RuleGroupId)),
	}.String())
}
func resolveWafRuleGroupRuleIds(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ruleGroup := resource.Item.(*types.RuleGroup)

	// Resolves rule group rules
	awsClient := meta.(*client.Client)
	service := awsClient.Services().Waf
	listActivatedRulesConfig := waf.ListActivatedRulesInRuleGroupInput{RuleGroupId: ruleGroup.RuleGroupId}
	var ruleIDs []string
	for {
		rules, err := service.ListActivatedRulesInRuleGroup(ctx, &listActivatedRulesConfig)
		if err != nil {
			return err
		}
		for _, rule := range rules.ActivatedRules {
			ruleIDs = append(ruleIDs, aws.ToString(rule.RuleId))
		}

		if aws.ToString(rules.NextMarker) == "" {
			break
		}
		listActivatedRulesConfig.NextMarker = rules.NextMarker
	}
	return resource.Set("rule_ids", ruleIDs)
}
func resolveWafRuleGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ruleGroup := resource.Item.(*types.RuleGroup)

	// Resolve tags for resource
	cl := meta.(*client.Client)
	service := cl.Services().Waf

	// Generate arn
	arnStr := arn.ARN{
		Partition: cl.Partition,
		Service:   "waf",
		Region:    "",
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("rulegroup/%s", aws.ToString(ruleGroup.RuleGroupId)),
	}.String()

	outputTags := make(map[string]*string)
	tagsConfig := waf.ListTagsForResourceInput{ResourceARN: aws.String(arnStr)}
	for {
		tags, err := service.ListTagsForResource(ctx, &tagsConfig, func(options *waf.Options) {
			// Set region to default global region
			options.Region = "us-east-1"
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
	return resource.Set("tags", outputTags)
}
