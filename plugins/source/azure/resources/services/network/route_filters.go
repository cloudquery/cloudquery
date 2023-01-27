package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RouteFilters() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_route_filters",
		Resolver:    fetchRouteFilters,
		Description: "https://learn.microsoft.com/en-us/rest/api/expressroute/route-filters/list?tabs=HTTP#routefilter",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_network_route_filters", client.Namespacemicrosoft_network),
		Transform:   transformers.TransformWithStruct(&armnetwork.RouteFilter{}),
		Columns: []schema.Column{
			client.SubscriptionID,
			client.IDColumn,
		},
	}
}

func fetchRouteFilters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewRouteFiltersClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
