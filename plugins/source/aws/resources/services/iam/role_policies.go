package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RolePolicies() *schema.Table {
	return &schema.Table{
		Name:                "aws_iam_role_policies",
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetRolePolicy.html`,
		Resolver:            fetchIamRolePolicies,
		PreResourceResolver: getRolePolicy,
		Transform:           transformers.TransformWithStruct(&iam.GetRolePolicyOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("iam"),
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
