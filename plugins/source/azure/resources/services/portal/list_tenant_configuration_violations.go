package portal

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ListTenantConfigurationViolations() *schema.Table {
	return &schema.Table{
		Name:                 "azure_portal_list_tenant_configuration_violations",
		Resolver:             fetchListTenantConfigurationViolations,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal@v0.5.0#Violation",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_portal_list_tenant_configuration_violations", client.Namespacemicrosoft_portal),
		Transform:            transformers.TransformWithStruct(&armportal.Violation{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
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
