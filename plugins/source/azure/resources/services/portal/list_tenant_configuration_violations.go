package portal

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ListTenantConfigurationViolations() *schema.Table {
	return &schema.Table{
		Name:      "azure_portal_list_tenant_configuration_violations",
		Resolver:  fetchListTenantConfigurationViolations,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_portal_list_tenant_configuration_violations", client.Namespacemicrosoft_portal),
		Transform: transformers.TransformWithStruct(&armportal.Violation{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchListTenantConfigurationViolations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armportal.NewListTenantConfigurationViolationsClient(cl.Creds, cl.Options)
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
