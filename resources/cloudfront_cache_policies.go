package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CloudfrontCachePolicies() *schema.Table {
	return &schema.Table{
		Name:         "aws_cloudfront_cache_policies",
		Resolver:     fetchCloudfrontCachePolicies,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "min_ttl",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("CachePolicy.CachePolicyConfig.MinTTL"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CachePolicy.CachePolicyConfig.Name"),
			},
			{
				Name:     "comment",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CachePolicy.CachePolicyConfig.Comment"),
			},
			{
				Name:     "default_ttl",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("CachePolicy.CachePolicyConfig.DefaultTTL"),
			},
			{
				Name:     "max_ttl",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("CachePolicy.CachePolicyConfig.MaxTTL"),
			},
			{
				Name:     "cookies_behavior",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.CookieBehavior"),
			},
			{
				Name:     "cookies",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.Cookies.Items"),
			},
			{
				Name:     "enable_accept_encoding_gzip",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingGzip"),
			},
			{
				Name:     "headers_behavior",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.HeaderBehavior"),
			},
			{
				Name:     "headers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.Headers.Items"),
			},
			{
				Name:     "query_strings_behavior",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStringBehavior"),
			},
			{
				Name:     "query_strings",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Items"),
			},
			{
				Name:     "enable_accept_encoding_brotli",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingBrotli"),
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CachePolicy.Id"),
			},
			{
				Name:     "last_modified_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CachePolicy.LastModifiedTime"),
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCloudfrontCachePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config cloudfront.ListCachePoliciesInput
	c := meta.(*client.Client)
	s := c.Services()
	svc := s.Cloudfront
	for {
		response, err := svc.ListCachePolicies(ctx, nil, func(options *cloudfront.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		if response.CachePolicyList != nil {
			res <- response.CachePolicyList.Items
		}

		if aws.ToString(response.CachePolicyList.NextMarker) == "" {
			break
		}
		config.Marker = response.CachePolicyList.NextMarker
	}
	return nil
}
