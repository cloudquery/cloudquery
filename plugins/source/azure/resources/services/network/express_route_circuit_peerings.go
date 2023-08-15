package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func expressRouteCircuitPeerings() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_express_route_circuit_peerings",
		Resolver:             fetchExpressRouteCircuitPeerings,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/expressroute/express-route-circuit-peerings/list?tabs=HTTP#expressroutecircuitpeering",
		Transform:            transformers.TransformWithStruct(&armnetwork.ExpressRouteCircuitPeering{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchExpressRouteCircuitPeerings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armnetwork.ExpressRouteCircuit)
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewExpressRouteCircuitPeeringsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*p.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(group, *p.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
