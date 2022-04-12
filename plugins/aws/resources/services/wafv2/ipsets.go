package wafv2

import (
	"context"
	"net"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen -config=ipsets.hcl -domain=wafv2 -resource=ipsets
func Ipsets() *schema.Table {
	return &schema.Table{
		Name:         "aws_wafv2_ipsets",
		Description:  "Contains one or more IP addresses or blocks of IP addresses specified in Classless Inter-Domain Routing (CIDR) notation",
		Resolver:     fetchWafv2Ipsets,
		Multiplex:    client.ServiceAccountRegionScopeMultiplexer("waf-regional"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name:        "addresses",
				Description: "Contains an array of strings that specify one or more IP addresses or blocks of IP addresses in Classless Inter-Domain Routing (CIDR) notation",
				Type:        schema.TypeCIDRArray,
				Resolver:    resolveIpsetAddresses,
			},
			{
				Name:        "ip_address_version",
				Description: "Specify IPV4 or IPV6.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IPAddressVersion"),
			},
			{
				Name:        "id",
				Description: "A unique identifier for the set",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the IP set",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "A description of the IP set that helps with identification.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveIpsetTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchWafv2Ipsets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafV2

	params := wafv2.ListIPSetsInput{
		Scope: cl.WAFScope,
		Limit: aws.Int32(100), // maximum value: https://docs.aws.amazon.com/waf/latest/APIReference/API_ListIPSets.html
	}
	for {
		result, err := svc.ListIPSets(ctx, &params, func(options *wafv2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, s := range result.IPSets {
			info, err := svc.GetIPSet(
				ctx,
				&wafv2.GetIPSetInput{
					Id:    s.Id,
					Name:  s.Name,
					Scope: cl.WAFScope,
				},
				func(options *wafv2.Options) { options.Region = cl.Region },
			)
			if err != nil {
				return diag.WrapError(err)
			}
			res <- info.IPSet
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return nil
}

func resolveIpsetAddresses(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(*types.IPSet)
	addrs := make([]*net.IPNet, 0, len(s.Addresses))
	for _, a := range s.Addresses {
		_, n, err := net.ParseCIDR(a)
		if err != nil {
			return diag.WrapError(err)
		}
		addrs = append(addrs, n)
	}
	return diag.WrapError(resource.Set(c.Name, addrs))
}

func resolveIpsetTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().WafV2
	s := resource.Item.(*types.IPSet)
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
