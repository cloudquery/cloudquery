package wafv2

import (
	"context"
	"encoding/json"
	"errors"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func RuleGroups() *schema.Table {
	tableName := "aws_wafv2_rule_groups"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/waf/latest/APIReference/API_RuleGroup.html`,
		Resolver:            fetchWafv2RuleGroups,
		PreResourceResolver: getRuleGroup,
		Transform:           transformers.TransformWithStruct(&types.RuleGroup{}),
		Multiplex:           client.ServiceAccountRegionScopeMultiplexer(tableName, "waf-regional"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveRuleGroupTags,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ARN"),
				PrimaryKey: true,
			},
			{
				Name:     "policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveWafv2ruleGroupPolicy,
			},
		},
	}
}

func fetchWafv2RuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafv2

	config := wafv2.ListRuleGroupsInput{Scope: cl.WAFScope}
	for {
		output, err := svc.ListRuleGroups(ctx, &config, func(o *wafv2.Options) {
			o.Region = cl.Region
		})
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
	cl := meta.(*client.Client)
	svc := cl.Services().Wafv2
	ruleGroupOutput := resource.Item.(types.RuleGroupSummary)

	// Get RuleGroup object
	ruleGroup, err := svc.GetRuleGroup(ctx, &wafv2.GetRuleGroupInput{
		Name:  ruleGroupOutput.Name,
		Id:    ruleGroupOutput.Id,
		Scope: cl.WAFScope,
	}, func(o *wafv2.Options) {
		o.Region = cl.Region
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
		tags, err := service.ListTagsForResource(ctx, &tagsConfig, func(o *wafv2.Options) {
			o.Region = cl.Region
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
func resolveWafv2ruleGroupPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ruleGroup := resource.Item.(*types.RuleGroup)

	cl := meta.(*client.Client)
	service := cl.Services().Wafv2

	// Resolve rule group policy
	policy, err := service.GetPermissionPolicy(ctx, &wafv2.GetPermissionPolicyInput{ResourceArn: ruleGroup.ARN}, func(o *wafv2.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		// we may get WAFNonexistentItemException error until SetPermissionPolicy is called on a rule group
		var e *types.WAFNonexistentItemException
		if errors.As(err, &e) {
			return resource.Set(c.Name, "null")
		}
		return err
	}
	var p map[string]any
	err = json.Unmarshal([]byte(*policy.Policy), &p)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, p)
}
