package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:                "aws_iam_users",
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_User.html`,
		Resolver:            fetchIamUsers,
		PreResourceResolver: getUser,
		Transform:           transformers.TransformWithStruct(&types.User{}),
		Multiplex:           client.AccountMultiplex,
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
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
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
		},
	}
}
