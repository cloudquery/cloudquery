package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func groupAttachedPolicies() *schema.Table {
	tableName := "aws_iam_group_attached_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachedPolicy.html`,
		Resolver:    fetchIamGroupAttachedPolicies,
		Transform:   transformers.TransformWithStruct(&types.AttachedPolicy{}, transformers.WithPrimaryKeys("PolicyArn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchIamGroupAttachedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(types.Group)
	svc := meta.(*client.Client).Services().Iam
	config := iam.ListAttachedGroupPoliciesInput{
		GroupName: p.GroupName,
	}
	paginator := iam.NewListAttachedGroupPoliciesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.AttachedPolicies
	}
	return nil
}
