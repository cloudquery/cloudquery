package s3

import (
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BucketGrants() *schema.Table {
	return &schema.Table{
		Name:        "aws_s3_bucket_grants",
		Description: `https://docs.aws.amazon.com/AmazonS3/latest/API/API_Grant.html`,
		Resolver:    fetchS3BucketGrants,
		Transform:   transformers.TransformWithStruct(&types.Grant{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:            "bucket_arn",
				Type:            schema.TypeString,
				Resolver:        schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "grantee_type",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("Grantee.Type"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "grantee_id",
				Type:            schema.TypeString,
				Resolver:        resolveBucketGranteeID,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}
