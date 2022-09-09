package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CloudfrontCachePolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudfront_cache_policies",
		Description: "Contains a cache policy.",
		Resolver:    fetchCloudfrontCachePolicies,
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARNWithAccount(client.CloudfrontService, func(resource *schema.Resource) ([]string, error) {
					return []string{"cache-policy", *resource.Item.(types.CachePolicySummary).CachePolicy.Id}, nil
				}),
			},
			{
				Name:     "cache_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CachePolicy"),
			},
			{
				Name:        "enable_accept_encoding_gzip",
				Description: "A flag that can affect whether the Accept-Encoding HTTP header is included in the cache key and included in requests that CloudFront sends to the origin",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingGzip"),
			},
			{
				Name:        "enable_accept_encoding_brotli",
				Description: "A flag that can affect whether the Accept-Encoding HTTP header is included in the cache key and included in requests that CloudFront sends to the origin",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("CachePolicy.CachePolicyConfig.ParametersInCacheKeyAndForwardedToOrigin.EnableAcceptEncodingBrotli"),
			},
			{
				Name:            "id",
				Description:     "The unique identifier for the cache policy",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("CachePolicy.Id"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchCloudfrontCachePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config cloudfront.ListCachePoliciesInput
	c := meta.(*client.Client)
	s := c.Services()
	svc := s.Cloudfront
	for {
		response, err := svc.ListCachePolicies(ctx, nil)
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
