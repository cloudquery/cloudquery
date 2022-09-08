package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/plugin-sdk/schema"
)


func Rules() *schema.Table {
	return &schema.Table{
		Name:         "aws_wafregional_rules",
		Description:  "A combination of identifiers for web requests that you want to allow, block, or count.",
		Resolver:     fetchWafregionalRules,
		Multiplex:    client.ServiceAccountRegionMultiplexer("waf-regional"),
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
				Description: "ARN of the rule.",
				Type:        schema.TypeString,
				Resolver:    resolveWafregionalRuleArn,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "tags",
				Description: "Rule tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveWafregionalRuleTags,
			},
			{
				Name:        "id",
				Description: "A unique identifier for a Rule",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RuleId"),
			},
			{
				Name:        "metric_name",
				Description: "A friendly name or description for the metrics for this Rule",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The friendly name or description for the Rule",
				Type:        schema.TypeString,
			},
			{
				Name:        "predicates",
				Description: "Contains one Predicate element for each ByteMatchSet, IPSet, or SqlInjectionMatchSet object that you want to include in a RateBasedRule.",
				Type:        schema.TypeJSON,
				Resolver: schema.PathResolver("Predicates"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchWafregionalRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafRegional
	var params wafregional.ListRulesInput
	for {
		result, err := svc.ListRules(ctx, &params, func(o *wafregional.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, r := range result.Rules {
			detail, err := svc.GetRule(
				ctx,
				&wafregional.GetRuleInput{RuleId: r.RuleId},
				func(o *wafregional.Options) {
					o.Region = cl.Region
				},
			)
			if err != nil {
				return err
			}
			if detail.Rule == nil {
				continue
			}
			res <- *detail.Rule
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return nil
}
func resolveWafregionalRuleArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return diag.WrapError(resource.Set(c.Name, ruleARN(meta, *resource.Item.(types.Rule).RuleId)))
}
func resolveWafregionalRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafRegional
	arn := ruleARN(meta, *resource.Item.(types.Rule).RuleId)
	params := wafregional.ListTagsForResourceInput{ResourceARN: &arn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			return err
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

func ruleARN(meta schema.ClientMeta, id string) string {
	cl := meta.(*client.Client)
	return cl.ARN(client.WAFRegional, "rule", id)
}
