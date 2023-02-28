package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/s3/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func bucketLifecycles() *schema.Table {
	return &schema.Table{
		Name:        "aws_s3_bucket_lifecycles",
		Description: `https://docs.aws.amazon.com/AmazonS3/latest/API/API_LifecycleRule.html`,
		Resolver:    fetchS3BucketLifecycles,
		Transform:   transformers.TransformWithStruct(&types.LifecycleRule{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "bucket_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchS3BucketLifecycles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.WrappedBucket)
	c := meta.(*client.Client)
	svc := c.Services().S3
	region := parent.Get("region").(*schema.Text)
	if region == nil {
		return nil
	}
	lifecycleOutput, err := svc.GetBucketLifecycleConfiguration(ctx, &s3.GetBucketLifecycleConfigurationInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = region.Str
	})
	if err != nil {
		if client.IsAWSError(err, "NoSuchLifecycleConfiguration") {
			return nil
		}
		return err
	}
	res <- lifecycleOutput.Rules
	return nil
}
