package wafregional

import (
	"context"
	"fmt"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func RateBasedRules() *schema.Table {
	tableName := "aws_wafregional_rate_based_rules"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_RateBasedRule.html`,
		Resolver:    fetchWafregionalRateBasedRules,
		Transform:   transformers.TransformWithStruct(&types.RateBasedRule{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "waf-regional"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveWafregionalRateBasedRuleArn,
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveWafregionalRateBasedRuleTags,
			},
		},
	}
}

func fetchWafregionalRateBasedRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafregional
	var params wafregional.ListRateBasedRulesInput
	for {
		result, err := svc.ListRateBasedRules(ctx, &params, func(o *wafregional.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, item := range result.Rules {
			detail, err := svc.GetRateBasedRule(
				ctx,
				&wafregional.GetRateBasedRuleInput{RuleId: item.RuleId},
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
func resolveWafregionalRateBasedRuleArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return resource.Set(c.Name, rateBasedRuleARN(meta, *resource.Item.(types.RateBasedRule).RuleId))
}
func resolveWafregionalRateBasedRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafregional
	arnStr := rateBasedRuleARN(meta, *resource.Item.(types.RateBasedRule).RuleId)
	params := wafregional.ListTagsForResourceInput{ResourceARN: &arnStr}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &params, func(o *wafregional.Options) {
			o.Region = cl.Region
		})
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
	return resource.Set(c.Name, tags)
}

func rateBasedRuleARN(meta schema.ClientMeta, id string) string {
	cl := meta.(*client.Client)
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.WAFRegional),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("ratebasedrule/%s", id),
	}.String()
}
