package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func delegatedServices() *schema.Table {
	tableName := "aws_organizations_delegated_services"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_DelegatedService.html`,
		Resolver:    fetchOrganizationsDelegatedServices,
		Transform:   transformers.TransformWithStruct(&types.DelegatedService{}, transformers.WithPrimaryKeys("ServicePrincipal")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}
func fetchOrganizationsDelegatedServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Organizations
	paginator := organizations.NewListDelegatedServicesForAccountPaginator(svc, &organizations.ListDelegatedServicesForAccountInput{
		AccountId: parent.Item.(types.Account).Id,
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.DelegatedServices
	}
	return nil
}
