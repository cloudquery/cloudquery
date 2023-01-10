package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Policies() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_policies",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_ManagedPolicyDetail.html`,
		Resolver:    fetchIamPolicies,
		Transform:   transformers.TransformWithStruct(&types.ManagedPolicyDetail{}),
		Multiplex:   client.AccountMultiplex,
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
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveIamPolicyTags,
			},
			{
				Name:     "policy_version_list",
				Type:     schema.TypeJSON,
				Resolver: resolveIamPolicyVersionList,
			},
		},
	}
}
