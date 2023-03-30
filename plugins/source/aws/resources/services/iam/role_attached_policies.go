package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func roleAttachedPolicies() *schema.Table {
	tableName := "aws_iam_role_attached_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachedPolicy.html`,
		Resolver:    fetchIamRoleAttachedPolicies,
		Transform:   transformers.TransformWithStruct(&types.AttachedPolicy{}, transformers.WithPrimaryKeys("PolicyName")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "role_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchIamRoleAttachedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*types.Role)
	svc := meta.(*client.Client).Services().Iam
	config := iam.ListAttachedRolePoliciesInput{
		RoleName: p.RoleName,
	}
	paginator := iam.NewListAttachedRolePoliciesPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.AttachedPolicies
	}
	return nil
}
