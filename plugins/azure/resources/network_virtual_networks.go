package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-08-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func NetworkVirtualNetworks() *schema.Table {
	return &schema.Table{
		Name:         "azure_network_virtual_networks",
		Resolver:     fetchNetworkVirtualNetworks,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "extended_location_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExtendedLocation.Name"),
			},
			{
				Name:     "extended_location_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExtendedLocation.Type"),
			},
			{
				Name:     "address_space_address_prefixes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("VirtualNetworkPropertiesFormat.AddressSpace.AddressPrefixes"),
			},
			{
				Name:     "dhcp_options_dns_servers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("VirtualNetworkPropertiesFormat.DhcpOptions.DNSServers"),
			},
			{
				Name:     "resource_guid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualNetworkPropertiesFormat.ResourceGUID"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualNetworkPropertiesFormat.ProvisioningState"),
			},
			{
				Name:     "enable_ddos_protection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("VirtualNetworkPropertiesFormat.EnableDdosProtection"),
			},
			{
				Name:     "enable_vm_protection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("VirtualNetworkPropertiesFormat.EnableVMProtection"),
			},
			{
				Name:     "ddos_protection_plan_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualNetworkPropertiesFormat.DdosProtectionPlan.ID"),
			},
			{
				Name:     "bgp_communities_virtual_network_community",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualNetworkPropertiesFormat.BgpCommunities.VirtualNetworkCommunity"),
			},
			{
				Name:     "bgp_communities_regional_community",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualNetworkPropertiesFormat.BgpCommunities.RegionalCommunity"),
			},
			{
				Name: "etag",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name: "location",
				Type: schema.TypeString,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "azure_networks_virtual_network_subnets",
				Resolver: fetchNetworksVirtualNetworkSubnets,
				Columns: []schema.Column{
					{
						Name:     "virtual_network_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "address_prefix",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.AddressPrefix"),
					},
					{
						Name:     "address_prefixes",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.AddressPrefixes"),
					},
					{
						Name:     "network_security_format_resource_guid",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.SecurityGroupPropertiesFormat.ResourceGUID"),
					},
					{
						Name:     "network_security_format_provisioning_state",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.SecurityGroupPropertiesFormat.ProvisioningState"),
					},
					{
						Name:     "network_security_group_etag",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.Etag"),
					},
					{
						Name:     "network_security_group_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.ID"),
					},
					{
						Name:     "network_security_group_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.Name"),
					},
					{
						Name:     "network_security_group_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.Type"),
					},
					{
						Name:     "network_security_group_location",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.Location"),
					},
					{
						Name:     "network_security_group_tags",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.Tags"),
					},
					{
						Name:     "route_table_disable_bgp_route_propagation",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.RouteTable.RouteTablePropertiesFormat.DisableBgpRoutePropagation"),
					},
					{
						Name:     "route_table_provisioning_state",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.RouteTable.RouteTablePropertiesFormat.ProvisioningState"),
					},
					{
						Name:     "route_table_resource_guid",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.RouteTable.RouteTablePropertiesFormat.ResourceGUID"),
					},
					{
						Name:     "route_table_etag",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.RouteTable.Etag"),
					},
					{
						Name:     "route_table_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.RouteTable.ID"),
					},
					{
						Name:     "route_table_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.RouteTable.Name"),
					},
					{
						Name:     "route_table_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.RouteTable.Type"),
					},
					{
						Name:     "route_table_location",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.RouteTable.Location"),
					},
					{
						Name:     "route_table_tags",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.RouteTable.Tags"),
					},
					{
						Name:     "nat_gateway_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.NatGateway.ID"),
					},
					{
						Name:     "purpose",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.Purpose"),
					},
					{
						Name:     "provisioning_state",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.ProvisioningState"),
					},
					{
						Name:     "private_endpoint_network_policies",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.PrivateEndpointNetworkPolicies"),
					},
					{
						Name:     "private_link_service_network_policies",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetPropertiesFormat.PrivateLinkServiceNetworkPolicies"),
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "etag",
						Type: schema.TypeString,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
				},
			},
			{
				Name:     "azure_networks_virtual_network_peerings",
				Resolver: fetchNetworksVirtualNetworkPeerings,
				Columns: []schema.Column{
					{
						Name:     "virtual_network_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "allow_virtual_network_access",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.AllowVirtualNetworkAccess"),
					},
					{
						Name:     "allow_forwarded_traffic",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.AllowForwardedTraffic"),
					},
					{
						Name:     "allow_gateway_transit",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.AllowGatewayTransit"),
					},
					{
						Name:     "use_remote_gateways",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.UseRemoteGateways"),
					},
					{
						Name:     "remote_virtual_network_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.RemoteVirtualNetwork.ID"),
					},
					{
						Name:     "remote_address_space_address_prefixes",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.RemoteAddressSpace.AddressPrefixes"),
					},
					{
						Name:     "remote_bgp_communities_virtual_network_community",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.RemoteBgpCommunities.VirtualNetworkCommunity"),
					},
					{
						Name:     "remote_bgp_communities_regional_community",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.RemoteBgpCommunities.RegionalCommunity"),
					},
					{
						Name:     "peering_state",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.PeeringState"),
					},
					{
						Name:     "provisioning_state",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.ProvisioningState"),
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "etag",
						Type: schema.TypeString,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
				},
			},
			{
				Name:     "azure_network_virtual_network_ip_allocations",
				Resolver: fetchNetworkVirtualNetworkIpAllocations,
				Columns: []schema.Column{
					{
						Name:     "virtual_network_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchNetworkVirtualNetworks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Network.VirtualNetworks
	response, err := svc.ListAll(ctx)
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
func fetchNetworksVirtualNetworkSubnets(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	vn := parent.Item.(network.VirtualNetwork)
	if vn.Subnets == nil {
		return nil
	}
	res <- *vn.Subnets
	return nil
}
func fetchNetworksVirtualNetworkPeerings(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	vn := parent.Item.(network.VirtualNetwork)
	if vn.VirtualNetworkPeerings == nil {
		return nil
	}
	res <- *vn.VirtualNetworkPeerings
	return nil
}
func fetchNetworkVirtualNetworkIpAllocations(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	vn := parent.Item.(network.VirtualNetwork)
	if vn.IPAllocations == nil {
		return nil
	}
	res <- *vn.IPAllocations
	return nil
}
