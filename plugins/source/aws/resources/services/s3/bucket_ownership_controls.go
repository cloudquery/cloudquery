package s3

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/s3/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func bucketOwnershipControls() *schema.Table {
	return &schema.Table{
		Name:        "aws_s3_bucket_ownership_controls",
		Description: `https://docs.aws.amazon.com/AmazonS3/latest/API/API_OwnershipControlsRule.html`,
		Resolver:    fetchBucketOwnershipControls,
		Transform:   transformers.TransformWithStruct(&types.OwnershipControlsRule{}, transformers.WithPrimaryKeyComponents("ObjectOwnership"), transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "bucket_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}
func fetchBucketOwnershipControls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.WrappedBucket)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceS3).S3
	getBucketOwnershipControlOutput, err := svc.GetBucketOwnershipControls(ctx, &s3.GetBucketOwnershipControlsInput{Bucket: r.Name}, func(o *s3.Options) {
		o.Region = r.Region
	})

	if err != nil {
		// If buckets ownership controls are not set it will return an error instead of empty result
		if client.IsAWSError(err, "OwnershipControlsNotFoundError") {
			return nil
		}

		if client.IgnoreAccessDeniedServiceDisabled(err) {
			return nil
		}

		return err
	}

	if getBucketOwnershipControlOutput == nil || getBucketOwnershipControlOutput.OwnershipControls == nil {
		return nil
	}

	res <- getBucketOwnershipControlOutput.OwnershipControls.Rules
	return nil
}
