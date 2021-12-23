package resources

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Wafv2RuleGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_wafv2_rule_groups",
		Description:  "A rule group defines a collection of rules to inspect and control web requests that you can use in a WebACL",
		Resolver:     fetchWafv2RuleGroups,
		Multiplex:    client.ServiceAccountRegionMultiplexer("waf-regional"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the entity.  ",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ARN"),
			},
			{
				Name:        "capacity",
				Description: "The web ACL capacity units (WCUs) required for this rule group",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "id",
				Description: "A unique identifier for the rule group",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
			},
			{
				Name:        "name",
				Description: "The name of the rule group",
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
				Name:        "custom_response_bodies",
				Description: "A map of custom response keys and content bodies",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "description",
				Description: "A description of the rule group that helps with identification.",
				Type:        schema.TypeString,
			},
			{
				Name:        "label_namespace",
				Description: "The label namespace prefix for this rule group",
				Type:        schema.TypeString,
			},
			{
				Name:        "rules",
				Description: "The Rule statements used to identify the web requests that you want to allow, block, or count",
				Type:        schema.TypeJSON,
				Resolver:    resolveWafv2ruleGroupRules,
			},
			{
				Name:        "available_labels",
				Description: "The labels that one or more rules in this rule group add to matching web ACLs.",
				Type:        schema.TypeStringArray,
				Resolver:    resolveWafv2AvailableLabels,
			},
			{
				Name:        "consumed_labels",
				Description: "The labels that one or more rules in this rule group add to matching web ACLs.",
				Type:        schema.TypeStringArray,
				Resolver:    resolveWafv2ConsumedLabels,
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

	// Dependent on the region select the right scope
	scope := types.ScopeRegional
	region := c.Region
	if region == strings.ToLower("global") {
		region = "us-east-1"
		scope = types.ScopeCloudfront
	}
	config := wafv2.ListRuleGroupsInput{Scope: scope}
	for {
		output, err := service.ListRuleGroups(ctx, &config, func(options *wafv2.Options) {
			options.Region = region
		})
		if err != nil {
			return err
		}

		// Get RuleGroup object
		for _, ruleGroupOutput := range output.RuleGroups {
			ruleGroup, err := service.GetRuleGroup(ctx, &wafv2.GetRuleGroupInput{
				Name:  ruleGroupOutput.Name,
				Id:    ruleGroupOutput.Id,
				Scope: scope,
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
func resolveWafv2ruleGroupPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ruleGroup, ok := resource.Item.(*types.RuleGroup)
	if !ok {
		return fmt.Errorf("not a RuleGroup instance: %#v", resource.Item)
	}

	cl := meta.(*client.Client)
	service := cl.Services().WafV2

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

	return resource.Set(c.Name, policy.Policy)
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
func resolveWafv2AvailableLabels(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ruleGroup, ok := resource.Item.(*types.RuleGroup)
	if !ok {
		return fmt.Errorf("not a RuleGroup instance: %#v", resource.Item)
	}
	labels := make([]string, len(ruleGroup.AvailableLabels))
	for i, l := range ruleGroup.AvailableLabels {
		labels[i] = *l.Name
	}
	return resource.Set(c.Name, labels)
}
func resolveWafv2ConsumedLabels(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ruleGroup, ok := resource.Item.(*types.RuleGroup)
	if !ok {
		return fmt.Errorf("not a RuleGroup instance: %#v", resource.Item)
	}
	labels := make([]string, len(ruleGroup.ConsumedLabels))
	for i, l := range ruleGroup.ConsumedLabels {
		labels[i] = *l.Name
	}
	return resource.Set(c.Name, labels)
}
