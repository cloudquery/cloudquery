package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"

func init() {
	tables := []Table{
		{
			Service:        "armnetwork",
			Name:           "express_route_gateways",
			Struct:         &armnetwork.ExpressRouteGateway{},
			ResponseStruct: &armnetwork.ExpressRouteCircuitsClientListAllResponse{},
			Client:         &armnetwork.ExpressRouteCircuitsClient{},
			ListFunc:       (&armnetwork.ExpressRouteGatewaysClient{}).ListBySubscription,
			NewFunc:        armnetwork.NewExpressRouteGatewaysClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteGateways",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Network)`,
			ExtraColumns:   DefaultExtraColumns,
			SkipFetch:      true,
		},
	}
	Tables = append(Tables, tables...)
}
