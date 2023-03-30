package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Roles() *schema.Table {
	tableName := "aws_iam_roles"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_Role.html`,
		Resolver:            fetchIamRoles,
		PreResourceResolver: getRole,
		Transform:           transformers.TransformWithStruct(&types.Role{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
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
			roleAttachedPolicies(),
			rolePolicies(),
			roleLastAccessedDetails(),
		},
	}
}
