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
		Resolver:            fetchIamRolePolicies,
		PreResourceResolver: getRolePolicy,
		Transform:           transformers.TransformWithStruct(&iam.GetRolePolicyOutput{}),
		Multiplex:           client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "role_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "policy_document",
				Type:     schema.TypeJSON,
				Resolver: resolveRolePoliciesPolicyDocument,
			},
		},
	}
}
