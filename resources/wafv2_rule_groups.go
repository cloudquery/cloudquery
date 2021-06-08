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

func Wafv2RuleGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_wafv2_rule_groups",
		Resolver:     fetchWafv2RuleGroups,
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveWafv2ruleGroupTags,
			},
			{
				Name:     "policy",
				Type:     schema.TypeJSON,
				Resolver: resolveWafv2ruleGroupPolicy,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
			},
			{
				Name: "capacity",
				Type: schema.TypeBigInt,
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
				Name:     "rules",
				Type:     schema.TypeJSON,
				Resolver: resolveWafv2ruleGroupRules,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_wafv2_rule_group_available_labels",
				Resolver: fetchWafv2RuleGroupAvailableLabels,
				Columns: []schema.Column{
					{
						Name:     "rule_group_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_wafv2_rule_group_consumed_labels",
				Resolver: fetchWafv2RuleGroupConsumedLabels,
				Columns: []schema.Column{
					{
						Name:     "rule_group_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchWafv2RuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().WafV2
	config := wafv2.ListRuleGroupsInput{}
	for {
		output, err := service.ListRuleGroups(ctx, &config, func(options *wafv2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		// Get RuleGroup object
		for _, ruleGroupOutput := range output.RuleGroups {
			ruleGroup, err := service.GetRuleGroup(ctx, &wafv2.GetRuleGroupInput{
				Name: ruleGroupOutput.Name,
				Id:   ruleGroupOutput.Id,
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

func resolveWafv2ruleGroupPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ruleGroup, ok := resource.Item.(*types.RuleGroup)
	if !ok {
		return fmt.Errorf("not a RuleGroup instance: %#v", resource.Item)
	}

	client := meta.(*client.Client)
	service := client.Services().WafV2

	// Resolve rule group policy
	policy, err := service.GetPermissionPolicy(ctx, &wafv2.GetPermissionPolicyInput{ResourceArn: ruleGroup.ARN}, func(options *wafv2.Options) {
		options.Region = client.Region
	})
	if err != nil {
		return err
	}
	if policy.Policy != nil {
		data, err := json.Marshal(policy.Policy)
		if err != nil {
			return nil
		}
		if err := resource.Set(c.Name, data); err != nil {
			return err
		}
	}
	return nil
}

func resolveWafv2ruleGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ruleGroup, ok := resource.Item.(*types.RuleGroup)
	if !ok {
		return fmt.Errorf("not a RuleGroup instance: %#v", resource.Item)
	}

	client := meta.(*client.Client)
	service := client.Services().WafV2

	// Resolve tags
	outputTags := make(map[string]*string)
	tagsConfig := wafv2.ListTagsForResourceInput{ResourceARN: ruleGroup.ARN}
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

func resolveWafv2ruleGroupRules(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ruleGroup, ok := resource.Item.(*types.RuleGroup)
	if !ok {
		return fmt.Errorf("not a RuleGroup instance: %#v", resource.Item)
	}
	if len(ruleGroup.Rules) == 0 {
		return nil
	}
	data, err := json.Marshal(ruleGroup.Rules)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

func fetchWafv2RuleGroupAvailableLabels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	ruleGroup, ok := parent.Item.(*types.RuleGroup)
	if !ok {
		return fmt.Errorf("not a RuleGroup instance: %#v", parent.Item)
	}
	res <- ruleGroup.AvailableLabels
	return nil
}

func fetchWafv2RuleGroupConsumedLabels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	ruleGroup, ok := parent.Item.(*types.RuleGroup)
	if !ok {
		return fmt.Errorf("not a RuleGroup instance: %#v", parent.Item)
	}
	res <- ruleGroup.ConsumedLabels
	return nil
}
