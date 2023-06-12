package organizations

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func organizationalUnitAccounts() *schema.Table {
	tableName := "aws_organizations_organizational_unit_accounts"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_ListAccountsForParent.html
The 'request_account_id' column is added to show from where the request was made.`,
		Resolver:  fetchAccountsForParent,
		Transform: transformers.TransformWithStruct(&types.Account{}, transformers.WithPrimaryKeys("Arn")),
		Columns: []schema.Column{
			{
				Name:       "request_account_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSAccount,
				PrimaryKey: true,
			},
			{
				Name:       "parent_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
		},
	}
}
func fetchAccountsForParent(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Organizations
	paginator := organizations.NewListAccountsForParentPaginator(svc, &organizations.ListAccountsForParentInput{
		ParentId: parent.Item.(includeParentOU).Id,
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *organizations.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Accounts
	}
	return nil
}
