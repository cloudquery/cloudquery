package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func expressRouteCircuitAuthorizations() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_express_route_circuit_authorizations",
		Resolver:             fetchExpressRouteCircuitAuthorizations,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/expressroute/express-route-circuit-authorizations/list?tabs=HTTP#expressroutecircuitauthorization",
		Transform:            transformers.TransformWithStruct(&armnetwork.ExpressRouteCircuitAuthorization{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchExpressRouteCircuitAuthorizations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armnetwork.ExpressRouteCircuit)
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewExpressRouteCircuitAuthorizationsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
