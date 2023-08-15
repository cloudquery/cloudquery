package s3

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/s3/models"
	"github.com/cloudquery/plugin-sdk/v4/scalar"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func bucketWebsites() *schema.Table {
	return &schema.Table{
		Name:        "aws_s3_bucket_websites",
		Description: `https://docs.aws.amazon.com/AmazonS3/latest/API/API_WebsiteConfiguration.html`,
		Resolver:    fetchS3BucketWebsites,
		Transform:   transformers.TransformWithStruct(&s3.GetBucketWebsiteOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "bucket_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
func fetchS3BucketWebsites(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.WrappedBucket)
	cl := meta.(*client.Client)
	svc := cl.Services().S3
	region := parent.Get("region").(*scalar.String)
	if region == nil {
		return nil
	}
	websiteOutput, err := svc.GetBucketWebsite(ctx, &s3.GetBucketWebsiteInput{Bucket: r.Name}, func(o *s3.Options) {
		o.Region = region.Value
	})
	if err != nil {
		if client.IsAWSError(err, "NoSuchBucket", "NoSuchWebsiteConfiguration") {
			return nil
		}
		return err
	}
	res <- websiteOutput
	return nil
}
