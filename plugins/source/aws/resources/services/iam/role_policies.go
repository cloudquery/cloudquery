package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RolePolicies() *schema.Table {
	tableName := "aws_iam_role_policies"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetRolePolicy.html`,
		Resolver:            fetchIamRolePolicies,
		PreResourceResolver: getRolePolicy,
		Transform:           client.TransformWithStruct(&iam.GetRolePolicyOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "role_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "policy_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "policy_document",
				Type:     schema.TypeJSON,
				Resolver: resolveRolePoliciesPolicyDocument,
			},
		},
	}
}
