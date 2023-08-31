package maintenance

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PublicMaintenanceConfigurations() *schema.Table {
	return &schema.Table{
		Name:                 "azure_maintenance_public_maintenance_configurations",
		Resolver:             fetchPublicMaintenanceConfigurations,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/maintenance/public-maintenance-configurations/list?tabs=HTTP#maintenanceconfiguration",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_maintenance_public_maintenance_configurations", client.Namespacemicrosoft_maintenance),
		Transform:            transformers.TransformWithStruct(&armmaintenance.Configuration{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchPublicMaintenanceConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armmaintenance.NewPublicMaintenanceConfigurationsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
