package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func UserPolicies() *schema.Table {
	return &schema.Table{
		Name:                "aws_iam_user_policies",
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetUserPolicy.html`,
		Resolver:            fetchIamUserPolicies,
		PreResourceResolver: getUserPolicy,
		Transform:           transformers.TransformWithStruct(&iam.GetUserPolicyOutput{}),
		Multiplex:           client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "user_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "policy_document",
				Type:     schema.TypeJSON,
				Resolver: resolveIamUserPolicyPolicyDocument,
			},
		},
	}
}
