package iam

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func roleAttachedPolicies() *schema.Table {
	tableName := "aws_iam_role_attached_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachedPolicy.html`,
		Resolver:    fetchIamRoleAttachedPolicies,
		Transform:   transformers.TransformWithStruct(&types.AttachedPolicy{}, transformers.WithPrimaryKeys("PolicyArn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "role_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchIamRoleAttachedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*types.Role)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	config := iam.ListAttachedRolePoliciesInput{
		RoleName: p.RoleName,
	}
	paginator := iam.NewListAttachedRolePoliciesPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.AttachedPolicies
	}
	return nil
}
