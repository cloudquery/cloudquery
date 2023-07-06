package lightsail

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Buckets() *schema.Table {
	tableName := "aws_lightsail_buckets"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Bucket.html`,
		Resolver:    fetchLightsailBuckets,
		Transform:   transformers.TransformWithStruct(&types.Bucket{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "able_to_update_bundle",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("AbleToUpdateBundle"),
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			bucketAccessKeys(),
		},
	}
}

func fetchLightsailBuckets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input lightsail.GetBucketsInput
	cl := meta.(*client.Client)
	svc := cl.Services().Lightsail
	// No paginator available
	for {
		response, err := svc.GetBuckets(ctx, &input, func(options *lightsail.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Buckets
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
