// Auto generated code - DO NOT EDIT.

package network

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ExpressRouteGateways() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_express_route_gateways",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network#ExpressRouteGateway`,
		Resolver:    fetchNetworkExpressRouteGateways,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "auto_scale_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AutoScaleConfiguration"),
			},
			{
				Name:     "express_route_connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExpressRouteConnections"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "virtual_hub",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VirtualHub"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}

func fetchNetworkExpressRouteGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.ExpressRouteGateways

	response, err := svc.ListBySubscription(ctx)
	if err != nil {
		return err
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value

	return nil
}
