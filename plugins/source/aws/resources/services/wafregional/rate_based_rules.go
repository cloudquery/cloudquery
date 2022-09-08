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


func RateBasedRules() *schema.Table {
	return &schema.Table{
		Name:         "aws_wafregional_rate_based_rules",
		Description:  "A combination of identifiers for web requests that you want to allow, block, or count, including rate limit.",
		Resolver:     fetchWafregionalRateBasedRules,
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
				Description: "ARN of the rate based rule.",
				Type:        schema.TypeString,
				Resolver:    resolveWafregionalRateBasedRuleArn,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "tags",
				Description: "Rule tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveWafregionalRateBasedRuleTags,
			},
			{
				Name:        "rate_key",
				Description: "The field that AWS WAF uses to determine if requests are likely arriving from single source and thus subject to rate monitoring",
				Type:        schema.TypeString,
			},
			{
				Name:        "rate_limit",
				Description: "The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period",
				Type:        schema.TypeInt,
			},
			{
				Name:        "id",
				Description: "A unique identifier for a RateBasedRule",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RuleId"),
			},
			{
				Name:        "metric_name",
				Description: "A friendly name or description for the metrics for a RateBasedRule",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "A friendly name or description for a RateBasedRule",
				Type:        schema.TypeString,
			},
			{
				Name:        "match_predicates",
				Description: "Contains one Predicate element for each ByteMatchSet, IPSet, or SqlInjectionMatchSet object that you want to include in a RateBasedRule.",
				Resolver: schema.PathResolver("MatchPredicates"),
				Type: schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchWafregionalRateBasedRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafRegional
	var params wafregional.ListRateBasedRulesInput
	for {
		result, err := svc.ListRateBasedRules(ctx, &params, func(o *wafregional.Options) { o.Region = cl.Region })
		if err != nil {
			return err
		}
		for _, item := range result.Rules {
			detail, err := svc.GetRateBasedRule(
				ctx,
				&wafregional.GetRateBasedRuleInput{RuleId: item.RuleId},
				func(o *wafregional.Options) { o.Region = cl.Region },
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
func resolveWafregionalRateBasedRuleArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return diag.WrapError(resource.Set(c.Name, rateBasedRuleARN(meta, *resource.Item.(types.RateBasedRule).RuleId)))
}
func resolveWafregionalRateBasedRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafRegional
	arn := rateBasedRuleARN(meta, *resource.Item.(types.RateBasedRule).RuleId)
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

func rateBasedRuleARN(meta schema.ClientMeta, id string) string {
	cl := meta.(*client.Client)
	return cl.ARN(client.WAFRegional, "ratebasedrule", id)
}
