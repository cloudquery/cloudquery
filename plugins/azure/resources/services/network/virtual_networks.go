package network

import (
	"context"
	"fmt"
	"net"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func NetworkVirtualNetworks() *schema.Table {
	return &schema.Table{
		Name:         "azure_network_virtual_networks",
		Description:  "Azure virtual network",
		Resolver:     fetchNetworkVirtualNetworks,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:          "extended_location_name",
				Description:   "The name of the extended location.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ExtendedLocation.Name"),
				IgnoreInTests: true,
			},
			{
				Name:          "extended_location_type",
				Description:   "The type of the extended location.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ExtendedLocation.Type"),
				IgnoreInTests: true,
			},
			{
				Name:        "address_space_address_prefixes",
				Description: "A list of address blocks reserved for this virtual network in CIDR notation.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("VirtualNetworkPropertiesFormat.AddressSpace.AddressPrefixes"),
			},
			{
				Name:          "dhcp_options_dns_servers",
				Description:   "The list of DNS servers IP addresses.",
				Type:          schema.TypeInetArray,
				Resolver:      resolveNetworkVirtualNetworksDhcpOptionsDnsServers,
				IgnoreInTests: true,
			},
			{
				Name:        "resource_guid",
				Description: "The resourceGuid property of the Virtual Network resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualNetworkPropertiesFormat.ResourceGUID"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the virtual network resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualNetworkPropertiesFormat.ProvisioningState"),
			},
			{
				Name:        "enable_ddos_protection",
				Description: "Indicates if DDoS protection is enabled for all the protected resources in the virtual network",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualNetworkPropertiesFormat.EnableDdosProtection"),
			},
			{
				Name:          "enable_vm_protection",
				Description:   "Indicates if VM protection is enabled for all the subnets in the virtual network.",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("VirtualNetworkPropertiesFormat.EnableVMProtection"),
				IgnoreInTests: true,
			},
			{
				Name:          "ddos_protection_plan_id",
				Description:   "Resource ID.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualNetworkPropertiesFormat.DdosProtectionPlan.ID"),
				IgnoreInTests: true,
			},
			{
				Name:          "bgp_communities_virtual_network_community",
				Description:   "The BGP community associated with the virtual network.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualNetworkPropertiesFormat.BgpCommunities.VirtualNetworkCommunity"),
				IgnoreInTests: true,
			},
			{
				Name:          "bgp_communities_regional_community",
				Description:   "The BGP community associated with the region of the virtual network.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("VirtualNetworkPropertiesFormat.BgpCommunities.RegionalCommunity"),
				IgnoreInTests: true,
			},
			{
				Name:          "ip_allocations",
				Description:   "Array of IpAllocation which reference this VNET.",
				Type:          schema.TypeStringArray,
				Resolver:      resolveNetworkVirtualNetworksIpAllocations,
				IgnoreInTests: true,
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "Resource ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource location.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_network_virtual_network_subnets",
				Description: "Azure virtual network subnet",
				Resolver:    fetchNetworksVirtualNetworkSubnets,
				Columns: []schema.Column{
					{
						Name:        "virtual_network_cq_id",
						Description: "Unique CloudQuery ID of azure_network_virtual_networks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "address_prefix",
						Description: "The address prefix for the subnet.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.AddressPrefix"),
					},
					{
						Name:          "address_prefixes",
						Description:   "List of address prefixes for the subnet.",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.AddressPrefixes"),
						IgnoreInTests: true,
					},
					{
						Name:          "security_group_properties_format_resource_guid",
						Description:   "The resource GUID property of the network security group resource.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.SecurityGroupPropertiesFormat.ResourceGUID"),
						IgnoreInTests: true,
					},
					{
						Name:        "security_group_properties_format_provisioning_state",
						Description: "The provisioning state of the network security group resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.SecurityGroupPropertiesFormat.ProvisioningState"),
					},
					{
						Name:          "network_security_group_etag",
						Description:   "A unique read-only string that changes whenever the resource is updated.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.Etag"),
						IgnoreInTests: true,
					},
					{
						Name:          "network_security_group_id",
						Description:   "Resource ID.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.ID"),
						IgnoreInTests: true,
					},
					{
						Name:          "network_security_group_name",
						Description:   "Resource name.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.Name"),
						IgnoreInTests: true,
					},
					{
						Name:          "network_security_group_type",
						Description:   "Resource type.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.Type"),
						IgnoreInTests: true,
					},
					{
						Name:          "network_security_group_location",
						Description:   "Resource location.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.Location"),
						IgnoreInTests: true,
					},
					{
						Name:          "network_security_group_tags",
						Description:   "Resource tags.",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.Tags"),
						IgnoreInTests: true,
					},
					{
						Name:          "route_table_disable_bgp_route_propagation",
						Description:   "Whether to disable the routes learned by BGP on that route table.",
						Type:          schema.TypeBool,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.RouteTable.RouteTablePropertiesFormat.DisableBgpRoutePropagation"),
						IgnoreInTests: true,
					},
					{
						Name:        "route_table_provisioning_state",
						Description: "The provisioning state of the route table resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.RouteTable.RouteTablePropertiesFormat.ProvisioningState"),
					},
					{
						Name:          "route_table_resource_guid",
						Description:   "The resource GUID property of the route table.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.RouteTable.RouteTablePropertiesFormat.ResourceGUID"),
						IgnoreInTests: true,
					},
					{
						Name:          "route_table_etag",
						Description:   "A unique read-only string that changes whenever the resource is updated.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.RouteTable.Etag"),
						IgnoreInTests: true,
					},
					{
						Name:          "route_table_id",
						Description:   "Resource ID.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.RouteTable.ID"),
						IgnoreInTests: true,
					},
					{
						Name:          "route_table_name",
						Description:   "Resource name.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.RouteTable.Name"),
						IgnoreInTests: true,
					},
					{
						Name:          "route_table_type",
						Description:   "Resource type.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.RouteTable.Type"),
						IgnoreInTests: true,
					},
					{
						Name:          "route_table_location",
						Description:   "Resource location.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.RouteTable.Location"),
						IgnoreInTests: true,
					},
					{
						Name:          "route_table_tags",
						Description:   "Resource tags.",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.RouteTable.Tags"),
						IgnoreInTests: true,
					},
					{
						Name:          "ip_configurations",
						Description:   "IPConfigurations - READ-ONLY; An array of references to the network interface IP configurations using subnet.",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.IPConfigurations"),
						IgnoreInTests: true,
					},
					{
						Name:          "private_endpoints",
						Description:   "PrivateEndpoints - READ-ONLY; An array of references to private endpoints.",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.PrivateEndpoints"),
						IgnoreInTests: true,
					},
					{
						Name:          "nat_gateway_id",
						Description:   "Resource ID.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SubnetPropertiesFormat.NatGateway.ID"),
						IgnoreInTests: true,
					},
					{
						Name:        "purpose",
						Description: "A read-only string identifying the intention of use for this subnet based on delegations and other user-defined properties.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.Purpose"),
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the subnet resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.ProvisioningState"),
					},
					{
						Name:        "private_endpoint_network_policies",
						Description: "Enable or Disable apply network policies on private end point in the subnet.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.PrivateEndpointNetworkPolicies"),
					},
					{
						Name:        "private_link_service_network_policies",
						Description: "Enable or Disable apply network policies on private link service in the subnet.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.PrivateLinkServiceNetworkPolicies"),
					},
					{
						Name:        "name",
						Description: "The name of the resource that is unique within a resource group",
						Type:        schema.TypeString,
					},
					{
						Name:        "etag",
						Description: "A unique read-only string that changes whenever the resource is updated.",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
			},
			{
				Name:          "azure_network_virtual_network_peerings",
				Description:   "Azure virtual network peering",
				Resolver:      fetchNetworksVirtualNetworkPeerings,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "virtual_network_cq_id",
						Description: "Unique CloudQuery ID of azure_network_virtual_networks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allow_virtual_network_access",
						Description: "Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.AllowVirtualNetworkAccess"),
					},
					{
						Name:        "allow_forwarded_traffic",
						Description: "Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.AllowForwardedTraffic"),
					},
					{
						Name:        "allow_gateway_transit",
						Description: "If gateway links can be used in remote virtual networking to link to this virtual network.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.AllowGatewayTransit"),
					},
					{
						Name:        "use_remote_gateways",
						Description: "If remote gateways can be used on this virtual network",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.UseRemoteGateways"),
					},
					{
						Name:        "remote_virtual_network_id",
						Description: "Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.RemoteVirtualNetwork.ID"),
					},
					{
						Name:        "remote_address_space_address_prefixes",
						Description: "A list of address blocks reserved for this virtual network in CIDR notation.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.RemoteAddressSpace.AddressPrefixes"),
					},
					{
						Name:          "remote_bgp_communities_virtual_network_community",
						Description:   "The BGP community associated with the virtual network.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.RemoteBgpCommunities.VirtualNetworkCommunity"),
						IgnoreInTests: true,
					},
					{
						Name:          "remote_bgp_communities_regional_community",
						Description:   "The BGP community associated with the region of the virtual network.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.RemoteBgpCommunities.RegionalCommunity"),
						IgnoreInTests: true,
					},
					{
						Name:        "peering_state",
						Description: "The status of the virtual network peering",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.PeeringState"),
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the virtual network peering resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.ProvisioningState"),
					},
					{
						Name:        "name",
						Description: "The name of the resource that is unique within a resource group",
						Type:        schema.TypeString,
					},
					{
						Name:        "etag",
						Description: "A unique read-only string that changes whenever the resource is updated.",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchNetworkVirtualNetworks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.VirtualNetworks
	response, err := svc.ListAll(ctx)
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func resolveNetworkVirtualNetworksDhcpOptionsDnsServers(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	vn, ok := resource.Item.(network.VirtualNetwork)
	if !ok {
		return fmt.Errorf("expected to have network.VirtualNetwork but got %T", resource.Item)
	}
	if vn.DhcpOptions == nil || vn.DhcpOptions.DNSServers == nil || len(*vn.DhcpOptions.DNSServers) == 0 {
		return nil
	}
	ips := make([]net.IP, 0, len(*vn.DhcpOptions.DNSServers))
	for _, ip := range *vn.DhcpOptions.DNSServers {
		i := net.ParseIP(ip)
		if i == nil {
			return fmt.Errorf("wrong format of IP: %s", ip)
		}
		ips = append(ips, i)
		net.ParseIP(ip)
	}
	return resource.Set(c.Name, ips)
}
func resolveNetworkVirtualNetworksIpAllocations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	vn, ok := resource.Item.(network.VirtualNetwork)
	if !ok {
		return fmt.Errorf("expected to have network.VirtualNetwork but got %T", resource.Item)
	}
	if vn.IPAllocations == nil {
		return nil
	}
	allocations := make([]string, 0, len(*vn.IPAllocations))
	for _, a := range *vn.IPAllocations {
		allocations = append(allocations, *a.ID)
	}
	return resource.Set(c.Name, allocations)
}
func fetchNetworksVirtualNetworkSubnets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	vn, ok := parent.Item.(network.VirtualNetwork)
	if !ok {
		return fmt.Errorf("expected to have network.VirtualNetwork but got %T", parent.Item)
	}
	if vn.Subnets == nil {
		return nil
	}
	res <- *vn.Subnets
	return nil
}

func fetchNetworksVirtualNetworkPeerings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	vn, ok := parent.Item.(network.VirtualNetwork)
	if !ok {
		return fmt.Errorf("expected to have network.VirtualNetwork but got %T", parent.Item)
	}
	if vn.VirtualNetworkPeerings == nil {
		return nil
	}
	res <- *vn.VirtualNetworkPeerings
	return nil
}
