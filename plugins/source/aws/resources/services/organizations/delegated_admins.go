package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DelegatedAdministrators() *schema.Table {
	tableName := "aws_organizations_delegated_administrators"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_Root.html`,
		Resolver:    fetchOrganizationsDelegatedAdmins,
		Transform:   transformers.TransformWithStruct(&types.DelegatedAdministrator{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "organizations"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}
func fetchOrganizationsDelegatedAdmins(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Organizations
	var input organizations.ListDelegatedAdministratorsInput
	paginator := organizations.NewListDelegatedAdministratorsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.DelegatedAdministrators
	}
	return nil
}
