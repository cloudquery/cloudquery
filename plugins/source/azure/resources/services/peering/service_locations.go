package peering

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ServiceLocations() *schema.Table {
	return &schema.Table{
		Name:      "azure_peering_service_locations",
		Resolver:  fetchServiceLocations,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_peering_service_locations", client.Namespacemicrosoft_peering),
		Transform: transformers.TransformWithStruct(&armpeering.ServiceLocation{}),
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

func fetchServiceLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armpeering.NewServiceLocationsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
