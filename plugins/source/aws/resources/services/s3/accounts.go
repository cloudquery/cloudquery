// Code generated by codegen; DO NOT EDIT.

package s3

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:      "aws_s3_accounts",
		Resolver:  fetchS3Accounts,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "block_public_acls",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("BlockPublicAcls"),
			},
			{
				Name:     "block_public_policy",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("BlockPublicPolicy"),
			},
			{
				Name:     "ignore_public_acls",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IgnorePublicAcls"),
			},
			{
				Name:     "restrict_public_buckets",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RestrictPublicBuckets"),
			},
		},
	}
}
