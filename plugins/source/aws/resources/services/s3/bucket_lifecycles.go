package s3

import (
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BucketLifecycles() *schema.Table {
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
