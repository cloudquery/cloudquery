package portal

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TenantConfigurations() *schema.Table {
	return &schema.Table{
		Name:        "azure_portal_tenant_configurations",
		Resolver:    fetchTenantConfigurations,
		Description: "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal@v0.5.0#Configuration",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_portal_tenant_configurations", client.Namespacemicrosoft_portal),
		Transform:   transformers.TransformWithStruct(&armportal.Configuration{}),
		Columns:     schema.ColumnList{client.SubscriptionID, client.IDColumn},
	}
}

func fetchTenantConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armportal.NewTenantConfigurationsClient(cl.Creds, cl.Options)
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
