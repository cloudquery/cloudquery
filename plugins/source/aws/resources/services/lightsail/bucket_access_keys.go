package lightsail

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func bucketAccessKeys() *schema.Table {
	tableName := "aws_lightsail_bucket_access_keys"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_AccessKey.html`,
		Resolver:    fetchLightsailBucketAccessKeys,
		Transform:   transformers.TransformWithStruct(&types.AccessKey{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "bucket_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchLightsailBucketAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Bucket)
	cl := meta.(*client.Client)
	svc := cl.Services().Lightsail
	input := lightsail.GetBucketAccessKeysInput{
		BucketName: r.Name,
	}
	response, err := svc.GetBucketAccessKeys(ctx, &input, func(options *lightsail.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- response.AccessKeys
	return nil
}
