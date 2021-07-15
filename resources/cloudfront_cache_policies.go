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
		Description:  "Contains a cache policy.",
		Resolver:     fetchCloudfrontCachePolicies,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "min_ttl",
				Description: "The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.MinTTL"),
			},
			{
				Name:        "name",
				Description: "A unique name to identify the cache policy",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.Name"),
			},
			{
				Name:        "comment",
				Description: "A comment to describe the cache policy",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.Comment"),
			},
			{
				Name:        "default_ttl",
				Description: "The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.DefaultTTL"),
			},
			{
				Name:        "max_ttl",
				Description: "The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.MaxTTL"),
			},
			{
				Name:        "cookies_behavior",
				Description: "Determines whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.CookieBehavior"),
			},
			{
				Name:        "cookies_quantity",
				Description: "The number of cookie names in the Items list",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.Cookies.Quantity"),
			},
			{
				Name:        "cookies",
				Description: "A list of cookie names",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.CookiesConfig.Cookies.Items"),
			},
			{
				Name:        "enable_accept_encoding_gzip",
				Description: "A flag that can affect whether the Accept-Encoding HTTP header is included in the cache key and included in requests that CloudFront sends to the origin",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingGzip"),
			},
			{
				Name:        "headers_behavior",
				Description: "Determines whether any HTTP headers are included in the cache key and automatically included in requests that CloudFront sends to the origin",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.HeaderBehavior"),
			},
			{
				Name:        "headers_quantity",
				Description: "The number of header names in the Items list",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.Headers.Quantity"),
			},
			{
				Name:        "headers",
				Description: "A list of HTTP header names",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.HeadersConfig.Headers.Items"),
			},
			{
				Name:        "query_strings_behavior",
				Description: "Determines whether any URL query strings in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStringBehavior"),
			},
			{
				Name:        "query_strings_quantity",
				Description: "The number of query string names in the Items list",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Quantity"),
			},
			{
				Name:        "query_strings",
				Description: "A list of query string names",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.QueryStringsConfig.QueryStrings.Items"),
			},
			{
				Name:        "enable_accept_encoding_brotli",
				Description: "A flag that can affect whether the Accept-Encoding HTTP header is included in the cache key and included in requests that CloudFront sends to the origin",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingBrotli"),
			},
			{
				Name:        "id",
				Description: "The unique identifier for the cache policy",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CachePolicy.Id"),
			},
			{
				Name:        "last_modified_time",
				Description: "The date and time when the cache policy was last modified",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("CachePolicy.LastModifiedTime"),
			},
			{
				Name:        "type",
				Description: "The type of cache policy, either managed (created by AWS) or custom (created in this AWS account)",
				Type:        schema.TypeString,
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
