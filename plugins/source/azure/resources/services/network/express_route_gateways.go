package network

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ExpressRouteGateways() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_express_route_gateways",
		Resolver:             fetchExpressRouteGateways,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/expressroute/express-route-gateways/list-by-subscription?tabs=HTTP#expressroutegateway",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_network_express_route_gateways", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armnetwork.ExpressRouteGateway{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}
