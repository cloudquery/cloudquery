package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen -config=rate_based_rules.hcl -domain=wafregional -resource=rate_based_rules
func RateBasedRules() *schema.Table {
	return &schema.Table{
		Name:         "aws_wafregional_rate_based_rules",
		Description:  "This is AWS WAF Classic documentation",
		Resolver:     fetchWafregionalRateBasedRules,
		Multiplex:    client.ServiceAccountRegionMultiplexer("waf-regional"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
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
				Description: "ARN of the rate based rule.",
				Type:        schema.TypeString,
				Resolver:    resolveRateBasedRuleARN,
			},
			{
				Name:        "tags",
				Description: "Rule tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRateBasedRuleTags,
			},
			{
				Name:        "rate_key",
				Description: "The field that AWS WAF uses to determine if requests are likely arriving from single source and thus subject to rate monitoring",
				Type:        schema.TypeString,
			},
			{
				Name:        "rate_limit",
				Description: "The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period",
				Type:        schema.TypeBigInt,
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
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_wafregional_rate_based_rule_match_predicates",
				Description: "This is AWS WAF Classic documentation",
				Resolver:    fetchWafregionalRateBasedRuleMatchPredicates,
				Columns: []schema.Column{
					{
						Name:        "rate_based_rule_cq_id",
						Description: "Unique CloudQuery ID of aws_wafregional_rate_based_rules table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "data_id",
						Description: "A unique identifier for a predicate in a Rule, such as ByteMatchSetId or IPSetId",
						Type:        schema.TypeString,
					},
					{
						Name:        "negated",
						Description: "Set Negated to False if you want AWS WAF to allow, block, or count requests based on the settings in the specified ByteMatchSet, IPSet, SqlInjectionMatchSet, XssMatchSet, RegexMatchSet, GeoMatchSet, or SizeConstraintSet",
						Type:        schema.TypeBool,
					},
					{
						Name:        "type",
						Description: "The type of predicate in a Rule, such as ByteMatch or IPSet.  This member is required.",
						Type:        schema.TypeString,
					},
				},
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
			return diag.WrapError(err)
		}
		for _, item := range result.Rules {
			detail, err := svc.GetRateBasedRule(
				ctx,
				&wafregional.GetRateBasedRuleInput{RuleId: item.RuleId},
				func(o *wafregional.Options) { o.Region = cl.Region },
			)
			if err != nil {
				return diag.WrapError(err)
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
func resolveRateBasedRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafRegional
	arn := rateBasedRuleARN(meta, *resource.Item.(types.RateBasedRule).RuleId)
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
func fetchWafregionalRateBasedRuleMatchPredicates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	rule := parent.Item.(types.RateBasedRule)
	res <- rule.MatchPredicates
	return nil
}
func resolveRateBasedRuleARN(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return diag.WrapError(resource.Set(c.Name, rateBasedRuleARN(meta, *resource.Item.(types.RateBasedRule).RuleId)))
}
func rateBasedRuleARN(meta schema.ClientMeta, id string) string {
	cl := meta.(*client.Client)
	return cl.ARN(client.WAFRegional, "ratebasedrule", id)
}
