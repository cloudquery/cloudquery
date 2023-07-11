package iam

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func groupAttachedPolicies() *schema.Table {
	tableName := "aws_iam_group_attached_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachedPolicy.html`,
		Resolver:    fetchIamGroupAttachedPolicies,
		Transform:   transformers.TransformWithStruct(&types.AttachedPolicy{}, transformers.WithPrimaryKeys("PolicyArn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "group_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchIamGroupAttachedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(types.Group)
	cl := meta.(*client.Client)
	svc := cl.Services().Iam
	config := iam.ListAttachedGroupPoliciesInput{
		GroupName: p.GroupName,
	}
	paginator := iam.NewListAttachedGroupPoliciesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.AttachedPolicies
	}
	return nil
}
