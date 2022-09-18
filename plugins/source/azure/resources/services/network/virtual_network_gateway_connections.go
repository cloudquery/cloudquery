// Auto generated code - DO NOT EDIT.

package network

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
)

func virtualNetworkGatewayConnections() *schema.Table {
	return &schema.Table{
		Name:     "azure_network_virtual_network_gateway_connections",
		Resolver: fetchNetworkVirtualNetworkGatewayConnections,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "network_virtual_network_gateway_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIDResolver,
			},
			{
				Name:     "authorization_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthorizationKey"),
			},
			{
				Name:     "virtual_network_gateway_1",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VirtualNetworkGateway1"),
			},
			{
				Name:     "virtual_network_gateway_2",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VirtualNetworkGateway2"),
			},
			{
				Name:     "local_network_gateway_2",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LocalNetworkGateway2"),
			},
			{
				Name:     "connection_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectionType"),
			},
			{
				Name:     "connection_protocol",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectionProtocol"),
			},
			{
				Name:     "routing_weight",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RoutingWeight"),
			},
			{
				Name:     "connection_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectionMode"),
			},
			{
				Name:     "shared_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SharedKey"),
			},
			{
				Name:     "connection_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectionStatus"),
			},
			{
				Name:     "tunnel_connection_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TunnelConnectionStatus"),
			},
			{
				Name:     "egress_bytes_transferred",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("EgressBytesTransferred"),
			},
			{
				Name:     "ingress_bytes_transferred",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("IngressBytesTransferred"),
			},
			{
				Name:     "peer",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Peer"),
			},
			{
				Name:     "enable_bgp",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableBgp"),
			},
			{
				Name:     "use_policy_based_traffic_selectors",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("UsePolicyBasedTrafficSelectors"),
			},
			{
				Name:     "ipsec_policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IpsecPolicies"),
			},
			{
				Name:     "traffic_selector_policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TrafficSelectorPolicies"),
			},
			{
				Name:     "resource_guid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceGUID"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "express_route_gateway_bypass",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ExpressRouteGatewayBypass"),
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

func fetchNetworkVirtualNetworkGatewayConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.VirtualNetworkGatewayConnections

	gateway := parent.Item.(network.VirtualNetworkGateway)
	resourceDetails, err := client.ParseResourceID(*gateway.ID)
	if err != nil {
		return err
	}
	response, err := svc.ListConnections(ctx, resourceDetails.ResourceGroup, *gateway.Name)

	if err != nil {
		return err
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
