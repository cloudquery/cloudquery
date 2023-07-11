package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func OriginAccessIdentities() *schema.Table {
	tableName := "aws_cloudfront_origin_access_identities"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ListCloudFrontOriginAccessIdentities.html`,
		Resolver:    fetchCloudfrontOriginAccessIdentities,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudfront"),
		Transform:   transformers.TransformWithStruct(&types.CloudFrontOriginAccessIdentitySummary{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}

func fetchCloudfrontOriginAccessIdentities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	s := cl.Services()
	svc := s.Cloudfront
	var config cloudfront.ListCloudFrontOriginAccessIdentitiesInput
	paginator := cloudfront.NewListCloudFrontOriginAccessIdentitiesPaginator(svc, &config)

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *cloudfront.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		if output.CloudFrontOriginAccessIdentityList != nil {
			res <- output.CloudFrontOriginAccessIdentityList.Items
		}
	}

	return nil
}
