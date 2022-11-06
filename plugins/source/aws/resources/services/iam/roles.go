// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Roles() *schema.Table {
	return &schema.Table{
		Name:                "aws_iam_roles",
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_Role.html`,
		Resolver:            fetchIamRoles,
		PreResourceResolver: getRole,
		Multiplex:           client.AccountMultiplex,
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
				Name:     "policies",
				Type:     schema.TypeJSON,
				Resolver: resolveIamRolePolicies,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "assume_role_policy_document",
				Type:     schema.TypeJSON,
				Resolver: resolveRolesAssumeRolePolicyDocument,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
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
				Name:     "role_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleName"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "max_session_duration",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxSessionDuration"),
			},
			{
				Name:     "permissions_boundary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PermissionsBoundary"),
			},
			{
				Name:     "role_last_used",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RoleLastUsed"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			RolePolicies(),
		},
	}
}
