// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:                "aws_iam_users",
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_User.html`,
		Resolver:            fetchIamUsers,
		PreResourceResolver: getUser,
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
				Name:     "create_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateDate"),
			},
			{
				Name:     "path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Path"),
			},
			{
				Name:     "user_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserName"),
			},
			{
				Name:     "password_last_used",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("PasswordLastUsed"),
			},
			{
				Name:     "permissions_boundary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PermissionsBoundary"),
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
