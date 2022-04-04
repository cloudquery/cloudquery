package network

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func NetworkExpressRouteGateways() *schema.Table {
	return &schema.Table{
		Name:         "azure_network_express_route_gateways",
		Description:  "Azure Network Express Route Gateways",
		Resolver:     fetchNetworkExpressRouteGateways,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "id",
				Description: "Resource ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "auto_scale_configuration_bound_max",
				Description: "Maximum number of scale units deployed for ExpressRoute gateway.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ExpressRouteGatewayProperties.AutoScaleConfiguration.Bounds.Max"),
			},
			{
				Name:        "auto_scale_configuration_bound_min",
				Description: "Minimum number of scale units deployed for ExpressRoute gateway.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ExpressRouteGatewayProperties.AutoScaleConfiguration.Bounds.Min"),
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource location.",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Resource name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the express route gateway resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRouteGatewayProperties.ProvisioningState"),
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "type",
				Description: "Resource type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "virtual_hub_id",
				Description: "The Virtual Hub where the ExpressRoute gateway is or will be deployed.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRouteGatewayProperties.VirtualHub.ID"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_network_express_route_connections",
				Description: "ExpressRouteConnection resource.",
				Resolver:    fetchNetworkExpressRouteConnections,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"express_route_gateway_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "express_route_gateway_cq_id",
						Description: "Unique CloudQuery ID of azure_network_express_route_gateways table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "authorization_key",
						Description: "Authorization key to establish the connection.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteConnectionProperties.AuthorizationKey"),
					},
					{
						Name:        "enable_internet_security",
						Description: "Enable internet security.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ExpressRouteConnectionProperties.EnableInternetSecurity"),
					},
					{
						Name:        "express_route_circuit_peering_id",
						Description: "The ID of the ExpressRoute circuit peering.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteConnectionProperties.ExpressRouteCircuitPeering.ID"),
					},
					{
						Name:        "name",
						Description: "Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the express route connection resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteConnectionProperties.ProvisioningState"),
					},
					{
						Name:        "routing_weight",
						Description: "The routing weight associated to the connection.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ExpressRouteConnectionProperties.RoutingWeight"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchNetworkExpressRouteGateways(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.ExpressRouteGateways
	response, err := svc.ListBySubscription(ctx)
	if err != nil {
		return diag.WrapError(err)
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value
	return nil
}
func fetchNetworkExpressRouteConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	erg, ok := parent.Item.(network.ExpressRouteGateway)
	if !ok {
		return fmt.Errorf("expected to have network.ExpressRouteGateway but got %T", parent.Item)
	}
	if erg.ExpressRouteGatewayProperties != nil && erg.ExpressRouteGatewayProperties.ExpressRouteConnections != nil {
		res <- *erg.ExpressRouteGatewayProperties.ExpressRouteConnections
	}
	return nil
}
