package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Users() *schema.Table {
	tableName := "aws_iam_users"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_User.html`,
		Resolver:            fetchIamUsers,
		PreResourceResolver: getUser,
		Transform:           client.TransformWithStruct(&types.User{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			client.DefaultAccountIDColumn(true),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			UserAccessKeys(),
			UserGroups(),
			UserAttachedPolicies(),
			UserPolicies(),
			SshPublicKeys(),
			SigningCertificates(),
			userLastAccessedDetails(),
		},
	}
}
