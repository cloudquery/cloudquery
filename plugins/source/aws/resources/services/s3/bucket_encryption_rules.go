// Code generated by codegen; DO NOT EDIT.

package s3

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func BucketEncryptionRules() *schema.Table {
	return &schema.Table{
		Name:        "aws_s3_bucket_encryption_rules",
		Description: `https://docs.aws.amazon.com/AmazonS3/latest/API/API_ServerSideEncryptionRule.html`,
		Resolver:    fetchS3BucketEncryptionRules,
		Multiplex:   client.AccountMultiplex,
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
			{
				Name:     "apply_server_side_encryption_by_default",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ApplyServerSideEncryptionByDefault"),
			},
			{
				Name:     "bucket_key_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("BucketKeyEnabled"),
			},
		},
	}
}
