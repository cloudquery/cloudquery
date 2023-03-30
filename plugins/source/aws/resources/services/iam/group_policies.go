package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GroupPolicies() *schema.Table {
	tableName := "aws_iam_group_policies"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetGroupPolicy.html`,
		Resolver:            fetchIamGroupPolicies,
		PreResourceResolver: getGroupPolicy,
		Transform:           client.TransformWithStruct(&iam.GetGroupPolicyOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "group_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "policy_document",
				Type:     schema.TypeJSON,
				Resolver: resolveIamGroupPolicyPolicyDocument,
			},
		},
	}
}
