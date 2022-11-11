package wafv2

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchWafv2RuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Wafv2

	config := wafv2.ListRuleGroupsInput{Scope: c.WAFScope}
	for {
		output, err := svc.ListRuleGroups(ctx, &config)
		if err != nil {
			return err
		}

		res <- output.RuleGroups

		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.NextMarker = output.NextMarker
	}
	return nil
}

func getRuleGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Wafv2
	ruleGroupOutput := resource.Item.(types.RuleGroupSummary)

	// Get RuleGroup object
	ruleGroup, err := svc.GetRuleGroup(ctx, &wafv2.GetRuleGroupInput{
		Name:  ruleGroupOutput.Name,
		Id:    ruleGroupOutput.Id,
		Scope: c.WAFScope,
	})
	if err != nil {
		return err
	}

	resource.Item = ruleGroup.RuleGroup
	return nil
}

func resolveRuleGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ruleGroup := resource.Item.(*types.RuleGroup)

	cl := meta.(*client.Client)
	service := cl.Services().Wafv2

	// Resolve tags
	outputTags := make(map[string]*string)
	tagsConfig := wafv2.ListTagsForResourceInput{ResourceARN: ruleGroup.ARN}
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
	return resource.Set(c.Name, outputTags)
}
func resolveWafv2ruleGroupPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ruleGroup := resource.Item.(*types.RuleGroup)

	cl := meta.(*client.Client)
	service := cl.Services().Wafv2

	// Resolve rule group policy
	policy, err := service.GetPermissionPolicy(ctx, &wafv2.GetPermissionPolicyInput{ResourceArn: ruleGroup.ARN}, func(options *wafv2.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		// we may get WAFNonexistentItemException error until SetPermissionPolicy is called on a rule group
		var e *types.WAFNonexistentItemException
		if errors.As(err, &e) {
			return resource.Set(c.Name, "null")
		}
		return err
	}
	var p map[string]interface{}
	err = json.Unmarshal([]byte(*policy.Policy), &p)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, p)
}
