package cloudfront

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func CachePolicies() *schema.Table {
	tableName := "aws_cloudfront_cache_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CachePolicySummary.html`,
		Resolver:    fetchCloudfrontCachePolicies,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudfront"),
		Transform:   transformers.TransformWithStruct(&types.CachePolicySummary{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("CachePolicy.Id"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchCloudfrontCachePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config cloudfront.ListCachePoliciesInput
	cl := meta.(*client.Client)
	s := cl.Services()
	svc := s.Cloudfront
	for {
		response, err := svc.ListCachePolicies(ctx, &config, func(options *cloudfront.Options) {
			options.Region = cl.Region
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
