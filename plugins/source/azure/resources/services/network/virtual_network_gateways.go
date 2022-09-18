// Auto generated code - DO NOT EDIT.

package network

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
)

func virtualNetworkGateways() *schema.Table {
	return &schema.Table{
		Name:     "azure_network_virtual_network_gateways",
		Resolver: fetchNetworkVirtualNetworkGateways,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cq_id_parent",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "extended_location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExtendedLocation"),
			},
			{
				Name:     "ip_configurations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IPConfigurations"),
			},
			{
				Name:     "gateway_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GatewayType"),
			},
			{
				Name:     "vpn_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpnType"),
			},
			{
				Name:     "vpn_gateway_generation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpnGatewayGeneration"),
			},
			{
				Name:     "enable_bgp",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableBgp"),
			},
			{
				Name:     "enable_private_ip_address",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnablePrivateIPAddress"),
			},
			{
				Name:     "active_active",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ActiveActive"),
			},
			{
				Name:     "gateway_default_site",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GatewayDefaultSite"),
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "vpn_client_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpnClientConfiguration"),
			},
			{
				Name:     "bgp_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BgpSettings"),
			},
			{
				Name:     "custom_routes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CustomRoutes"),
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
				Name:     "enable_dns_forwarding",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableDNSForwarding"),
			},
			{
				Name:     "inbound_dns_forwarding_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InboundDNSForwardingEndpoint"),
			},
			{
				Name:     "v_net_extended_location_resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VNetExtendedLocationResourceID"),
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

		Relations: []*schema.Table{
			virtualNetworkGatewayConnections(),
		},
	}
}

func fetchNetworkVirtualNetworkGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.VirtualNetworkGateways

	network := parent.Item.(network.VirtualNetwork)
	resourceDetails, err := client.ParseResourceID(*network.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	response, err := svc.List(ctx, resourceDetails.ResourceGroup)

	if err != nil {
		return errors.WithStack(err)
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
