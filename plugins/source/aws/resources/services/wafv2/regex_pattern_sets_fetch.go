package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchWafv2RegexPatternSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafv2

	params := wafv2.ListRegexPatternSetsInput{
		Scope: cl.WAFScope,
		Limit: aws.Int32(100), // maximum value: https://docs.aws.amazon.com/waf/latest/APIReference/API_ListRegexPatternSets.html
	}
	for {
		result, err := svc.ListRegexPatternSets(ctx, &params)
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
		func(options *wafv2.Options) {
			options.Region = cl.Region
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
		result, err := svc.ListTagsForResource(ctx, &params, func(options *wafv2.Options) {
			options.Region = cl.Region
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
