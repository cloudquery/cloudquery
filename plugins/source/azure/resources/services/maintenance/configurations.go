package maintenance

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Configurations() *schema.Table {
	return &schema.Table{
		Name:        "azure_maintenance_configurations",
		Resolver:    fetchConfigurations,
		Description: "https://learn.microsoft.com/en-us/rest/api/maintenance/maintenance-configurations/list?tabs=HTTP#maintenanceconfiguration",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_maintenance_configurations", client.Namespacemicrosoft_maintenance),
		Transform:   transformers.TransformWithStruct(&armmaintenance.Configuration{}),
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

func fetchConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armmaintenance.NewConfigurationsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
