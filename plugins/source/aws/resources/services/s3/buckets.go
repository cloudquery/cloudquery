package s3

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/s3/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Buckets() *schema.Table {
	return &schema.Table{
		Name:      "aws_s3_buckets",
		Resolver:  fetchS3Buckets,
		Transform: transformers.TransformWithStruct(&models.WrappedBucket{}),
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveBucketARN(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			BucketEncryptionRules(),
			BucketLifecycles(),
			BucketGrants(),
			BucketCorsRules(),
		},
	}
}
