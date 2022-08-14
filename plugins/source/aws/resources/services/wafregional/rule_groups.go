package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource rule_groups --config rule_groups.hcl --output .
func RuleGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_wafregional_rule_groups",
		Description:  "A collection of predefined rules that you can add to a web ACL.",
		Resolver:     fetchWafregionalRuleGroups,
		Multiplex:    client.ServiceAccountRegionMultiplexer("waf-regional"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "id"}},
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
				Name:        "arn",
				Description: "ARN of the rule group.",
				Type:        schema.TypeString,
				Resolver:    resolveWafregionalRuleGroupArn,
			},
			{
				Name:        "tags",
				Description: "Rule group tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveWafregionalRuleGroupTags,
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

func fetchWafregionalRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafRegional
	var params wafregional.ListRuleGroupsInput
	for {
		result, err := svc.ListRuleGroups(ctx, &params, func(o *wafregional.Options) { o.Region = cl.Region })
		if err != nil {
			return diag.WrapError(err)
		}
		for _, g := range result.RuleGroups {
			detail, err := svc.GetRuleGroup(
				ctx,
				&wafregional.GetRuleGroupInput{RuleGroupId: g.RuleGroupId},
				func(o *wafregional.Options) { o.Region = cl.Region },
			)
			if err != nil {
				return diag.WrapError(err)
			}
			if detail.RuleGroup == nil {
				continue
			}
			res <- *detail.RuleGroup
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return nil
}
func resolveWafregionalRuleGroupArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return diag.WrapError(resource.Set(c.Name, ruleGroupARN(meta, *resource.Item.(types.RuleGroup).RuleGroupId)))
}
func resolveWafregionalRuleGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafRegional
	arn := ruleGroupARN(meta, *resource.Item.(types.RuleGroup).RuleGroupId)
	params := wafregional.ListTagsForResourceInput{ResourceARN: &arn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			return diag.WrapError(err)
		}
		if result.TagInfoForResource != nil {
			client.TagsIntoMap(result.TagInfoForResource.TagList, tags)
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return diag.WrapError(resource.Set(c.Name, tags))
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func ruleGroupARN(meta schema.ClientMeta, id string) string {
	cl := meta.(*client.Client)
	return cl.ARN(client.WAFRegional, "rulegroup", id)
}
