package wafv2

import (
	"context"
	"net"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Ipsets() *schema.Table {
	tableName := "aws_wafv2_ipsets"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/waf/latest/APIReference/API_IPSet.html`,
		Resolver:            fetchWafv2Ipsets,
		Transform:           transformers.TransformWithStruct(&types.IPSet{}),
		PreResourceResolver: getIpset,
		Multiplex:           client.ServiceAccountRegionScopeMultiplexer(tableName, "waf-regional"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "addresses",
				Type:     arrow.ListOf(sdkTypes.ExtensionTypes.Inet),
				Resolver: resolveIpsetAddresses,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveIpsetTags,
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

func fetchWafv2Ipsets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafv2

	params := wafv2.ListIPSetsInput{
		Scope: cl.WAFScope,
		Limit: aws.Int32(100), // maximum value: https://docs.aws.amazon.com/waf/latest/APIReference/API_ListIPSets.html
	}
	for {
		result, err := svc.ListIPSets(ctx, &params, func(o *wafv2.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- result.IPSets

		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return nil
}

func getIpset(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafv2
	s := resource.Item.(types.IPSetSummary)
	input := &wafv2.GetIPSetInput{
		Id: s.Id, Name: s.Name, Scope: cl.WAFScope,
	}
	info, err := svc.GetIPSet(ctx, input, func(o *wafv2.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = info.IPSet
	return nil
}

func resolveIpsetAddresses(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(*types.IPSet)
	addrs := make([]*net.IPNet, 0, len(s.Addresses))
	for _, a := range s.Addresses {
		_, n, err := net.ParseCIDR(a)
		if err != nil {
			return err
		}
		addrs = append(addrs, n)
	}
	return resource.Set(c.Name, addrs)
}

func resolveIpsetTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafv2
	s := resource.Item.(*types.IPSet)
	var tagList []types.Tag
	params := wafv2.ListTagsForResourceInput{ResourceARN: s.ARN}

	for {
		result, err := svc.ListTagsForResource(ctx, &params, func(o *wafv2.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		tagList = append(tagList, result.TagInfoForResource.TagList...)
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return resource.Set(c.Name, client.TagsToMap(tagList))
}
