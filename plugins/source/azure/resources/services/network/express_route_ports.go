package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ExpressRoutePorts() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_express_route_ports",
		Resolver:             fetchExpressRoutePorts,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/expressroute/express-route-ports/list?tabs=HTTP#expressrouteport",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_network_express_route_ports", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armnetwork.ExpressRoutePort{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchExpressRoutePorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewExpressRoutePortsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
