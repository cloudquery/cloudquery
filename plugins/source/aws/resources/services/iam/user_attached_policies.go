package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func userAttachedPolicies() *schema.Table {
	tableName := "aws_iam_user_attached_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachedPolicy.html`,
		Resolver:    fetchIamUserAttachedPolicies,
		Transform:   transformers.TransformWithStruct(&types.AttachedPolicy{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "user_arn",
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
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchIamUserAttachedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*types.User)
	svc := meta.(*client.Client).Services().Iam
	config := iam.ListAttachedUserPoliciesInput{
		UserName: p.UserName,
	}
	paginator := iam.NewListAttachedUserPoliciesPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.AttachedPolicies
	}
	return nil
}
