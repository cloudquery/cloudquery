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

//go:generate cq-gen -config=rules.hcl -domain=wafregional -resource=rules
func Rules() *schema.Table {
	return &schema.Table{
		Name:         "aws_wafregional_rules",
		Description:  "This is AWS WAF Classic documentation",
		Resolver:     fetchWafregionalRules,
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
				Description: "ARN of the rule.",
				Type:        schema.TypeString,
				Resolver:    resolveRuleARN,
			},
			{
				Name:        "tags",
				Description: "Rule tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRuleTags,
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
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_wafregional_rule_predicates",
				Description: "This is AWS WAF Classic documentation",
				Resolver:    fetchWafregionalRulePredicates,
				Columns: []schema.Column{
					{
						Name:        "rule_cq_id",
						Description: "Unique CloudQuery ID of aws_wafregional_rules table (FK)",
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

func fetchWafregionalRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafRegional
	var params wafregional.ListRulesInput
	for {
		result, err := svc.ListRules(ctx, &params, func(o *wafregional.Options) { o.Region = cl.Region })
		if err != nil {
			return diag.WrapError(err)
		}
		for _, r := range result.Rules {
			detail, err := svc.GetRule(
				ctx,
				&wafregional.GetRuleInput{RuleId: r.RuleId},
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
func fetchWafregionalRulePredicates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.Rule)
	res <- r.Predicates
	return nil
}
func resolveRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafRegional
	arn := ruleARN(meta, *resource.Item.(types.Rule).RuleId)
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
func resolveRuleARN(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return diag.WrapError(resource.Set(c.Name, ruleARN(meta, *resource.Item.(types.Rule).RuleId)))
}
func ruleARN(meta schema.ClientMeta, id string) string {
	cl := meta.(*client.Client)
	return cl.ARN(client.WAFRegional, "rule", id)
}
