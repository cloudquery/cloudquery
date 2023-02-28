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

func bucketCorsRules() *schema.Table {
	return &schema.Table{
		Name:        "aws_s3_bucket_cors_rules",
		Description: `https://docs.aws.amazon.com/AmazonS3/latest/API/API_CORSRule.html`,
		Resolver:    fetchS3BucketCorsRules,
		Transform:   transformers.TransformWithStruct(&types.CORSRule{}),
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
func fetchS3BucketCorsRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.WrappedBucket)
	c := meta.(*client.Client)
	svc := c.Services().S3
	region := parent.Get("region").(*schema.Text)
	if region == nil {
		return nil
	}
	corsOutput, err := svc.GetBucketCors(ctx, &s3.GetBucketCorsInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = region.Str
	})
	if err != nil {
		if client.IsAWSError(err, "NoSuchCORSConfiguration", "NoSuchBucket") {
			return nil
		}
		return err
	}
	if corsOutput != nil {
		res <- corsOutput.CORSRules
	}
	return nil
}
