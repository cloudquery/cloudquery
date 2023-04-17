package subscription

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Tenants() *schema.Table {
	return &schema.Table{
		Name:        "azure_subscription_tenants",
		Resolver:    fetchTenants,
		Description: "https://learn.microsoft.com/en-us/rest/api/resources/tenants/list?tabs=HTTP#tenantiddescription",
		Transform:   transformers.TransformWithStruct(&armsubscription.TenantIDDescription{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{},
	}
}

func fetchTenants(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsubscription.NewTenantsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
