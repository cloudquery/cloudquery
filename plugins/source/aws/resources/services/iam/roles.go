package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Roles() *schema.Table {
	return &schema.Table{
		Name:                "aws_iam_roles",
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_Role.html`,
		Resolver:            fetchIamRoles,
		PreResourceResolver: getRole,
		Transform:           transformers.TransformWithStruct(&types.Role{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
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
