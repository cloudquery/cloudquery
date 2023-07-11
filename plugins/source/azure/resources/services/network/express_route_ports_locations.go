package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ExpressRoutePortsLocations() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_express_route_ports_locations",
		Resolver:             fetchExpressRoutePortsLocations,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/expressroute/express-route-ports-locations/list?tabs=HTTP#expressrouteportslocation",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_network_express_route_ports_locations", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armnetwork.ExpressRoutePortsLocation{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchExpressRoutePortsLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewExpressRoutePortsLocationsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
