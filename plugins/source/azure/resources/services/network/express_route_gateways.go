package network

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ExpressRouteGateways() *schema.Table {
	return &schema.Table{
		Name:      "azure_network_express_route_gateways",
		Resolver:  fetchExpressRouteGateways,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_network_express_route_gateways", client.Namespacemicrosoft_network),
		Transform: transformers.TransformWithStruct(&armnetwork.ExpressRouteGateway{}),
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
