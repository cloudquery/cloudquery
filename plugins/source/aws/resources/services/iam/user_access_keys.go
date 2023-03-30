package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func UserAccessKeys() *schema.Table {
	tableName := "aws_iam_user_access_keys"
	return &schema.Table{
		Name:                 tableName,
		Description:          `https://docs.aws.amazon.com/IAM/latest/APIReference/API_AccessKeyMetadata.html`,
		Resolver:             fetchIamUserAccessKeys,
		PostResourceResolver: postIamUserAccessKeyResolver,
		Transform:            client.TransformWithStruct(&models.AccessKeyWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:            client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "user_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "access_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccessKeyId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name: "last_used",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "last_used_service_name",
				Type: schema.TypeString,
			},
		},
	}
}
