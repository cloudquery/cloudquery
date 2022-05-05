package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen -config=regex_pattern_sets.hcl -domain=wafv2 -resource=regex_pattern_sets
func RegexPatternSets() *schema.Table {
	return &schema.Table{
		Name:         "aws_wafv2_regex_pattern_sets",
		Description:  "Contains one or more regular expressions",
		Resolver:     fetchWafv2RegexPatternSets,
		Multiplex:    client.ServiceAccountRegionScopeMultiplexer("waf-regional"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionScopeFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "scope",
				Description: "Specifies whether this is for an Amazon CloudFront distribution or for a regional application.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveWAFScope,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the entity.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ARN"),
			},
			{
				Name:        "description",
				Description: "A description of the set that helps with identification.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "A unique identifier for the set",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the set",
				Type:        schema.TypeString,
			},
			{
				Name:        "regular_expression_list",
				Description: "The regular expression patterns in the set.",
				Type:        schema.TypeStringArray,
				Resolver:    resolveRegexPatternSetsRegularExpressionList,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveRegexPatternSetTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchWafv2RegexPatternSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafV2

	params := wafv2.ListRegexPatternSetsInput{
		Scope: cl.WAFScope,
		Limit: aws.Int32(100), // maximum value: https://docs.aws.amazon.com/waf/latest/APIReference/API_ListRegexPatternSets.html
	}
	for {
		result, err := svc.ListRegexPatternSets(ctx, &params, func(options *wafv2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, s := range result.RegexPatternSets {
			info, err := svc.GetRegexPatternSet(
				ctx,
				&wafv2.GetRegexPatternSetInput{
					Id:    s.Id,
					Name:  s.Name,
					Scope: cl.WAFScope,
				},
				func(options *wafv2.Options) { options.Region = cl.Region },
			)
			if err != nil {
				return diag.WrapError(err)
			}
			res <- info.RegexPatternSet
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return nil
}

func resolveRegexPatternSetsRegularExpressionList(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(*types.RegexPatternSet)
	items := make([]string, len(s.RegularExpressionList))
	for i, v := range s.RegularExpressionList {
		if v.RegexString != nil {
			items[i] = *v.RegexString
		}
	}
	return diag.WrapError(resource.Set(c.Name, items))
}

func resolveRegexPatternSetTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafV2
	s := resource.Item.(*types.RegexPatternSet)
	tags := make(map[string]string)
	params := wafv2.ListTagsForResourceInput{ResourceARN: s.ARN}
	for {
		result, err := svc.ListTagsForResource(ctx, &params, func(options *wafv2.Options) { options.Region = cl.Region })
		if err != nil {
			return diag.WrapError(err)
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
	return diag.WrapError(resource.Set(c.Name, tags))
}
