package wafv2

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func RegexPatternSets() *schema.Table {
	tableName := "aws_wafv2_regex_pattern_sets"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/waf/latest/APIReference/API_RegexPatternSet.html`,
		Resolver:            fetchWafv2RegexPatternSets,
		PreResourceResolver: getRegexPatternSet,
		Transform:           transformers.TransformWithStruct(&types.RegexPatternSet{}),
		Multiplex:           client.ServiceAccountRegionScopeMultiplexer(tableName, "waf-regional"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveRegexPatternSetTags,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ARN"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchWafv2RegexPatternSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafv2

	params := wafv2.ListRegexPatternSetsInput{
		Scope: cl.WAFScope,
		Limit: aws.Int32(100), // maximum value: https://docs.aws.amazon.com/waf/latest/APIReference/API_ListRegexPatternSets.html
	}
	for {
		result, err := svc.ListRegexPatternSets(ctx, &params, func(o *wafv2.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- result.RegexPatternSets

		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return nil
}

func getRegexPatternSet(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafv2
	s := resource.Item.(types.RegexPatternSetSummary)

	info, err := svc.GetRegexPatternSet(
		ctx,
		&wafv2.GetRegexPatternSetInput{
			Id:    s.Id,
			Name:  s.Name,
			Scope: cl.WAFScope,
		},
		func(o *wafv2.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}

	resource.Item = info.RegexPatternSet
	return nil
}

func resolveRegexPatternSetTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafv2
	s := resource.Item.(*types.RegexPatternSet)
	tags := make(map[string]string)
	params := wafv2.ListTagsForResourceInput{ResourceARN: s.ARN}
	for {
		result, err := svc.ListTagsForResource(ctx, &params, func(o *wafv2.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		if result != nil || result.TagInfoForResource != nil {
			for _, t := range result.TagInfoForResource.TagList {
				tags[aws.ToString(t.Key)] = aws.ToString(t.Value)
			}
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return resource.Set(c.Name, tags)
}
