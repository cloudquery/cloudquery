package wafv2

import (
	"context"
	"net"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchWafv2Ipsets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafv2

	params := wafv2.ListIPSetsInput{
		Scope: cl.WAFScope,
		Limit: aws.Int32(100), // maximum value: https://docs.aws.amazon.com/waf/latest/APIReference/API_ListIPSets.html
	}
	for {
		result, err := svc.ListIPSets(ctx, &params)
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

	info, err := svc.GetIPSet(
		ctx,
		&wafv2.GetIPSetInput{
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
