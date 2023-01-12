package s3

import (
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BucketEncryptionRules() *schema.Table {
	return &schema.Table{
		Name:        "aws_s3_bucket_encryption_rules",
		Description: `https://docs.aws.amazon.com/AmazonS3/latest/API/API_ServerSideEncryptionRule.html`,
		Resolver:    fetchS3BucketEncryptionRules,
		Transform:   transformers.TransformWithStruct(&types.ServerSideEncryptionRule{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "bucket_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
