package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
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
				Name:        "extended_location_name",
				Description: "The name of the extended location",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExtendedLocation.Name"),
			},
			{
				Name:        "extended_location_type",
				Description: "The type of the extended location",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExtendedLocation.Type"),
			},
			{
				Name:        "address_space_address_prefixes",
				Description: "A list of address blocks reserved for this virtual network in CIDR notation",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("VirtualNetworkPropertiesFormat.AddressSpace.AddressPrefixes"),
			},
			{
				Name:        "dhcp_options_dns_servers",
				Description: "The list of DNS servers IP addresses",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("VirtualNetworkPropertiesFormat.DhcpOptions.DNSServers"),
			},
			{
				Name:        "resource_guid",
				Description: "The resourceGuid property of the Virtual Network resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualNetworkPropertiesFormat.ResourceGUID"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the virtual network resource Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualNetworkPropertiesFormat.ProvisioningState"),
			},
			{
				Name:        "enable_ddos_protection",
				Description: "Indicates if DDoS protection is enabled for all the protected resources in the virtual network It requires a DDoS protection plan associated with the resource",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualNetworkPropertiesFormat.EnableDdosProtection"),
			},
			{
				Name:        "enable_vm_protection",
				Description: "Indicates if VM protection is enabled for all the subnets in the virtual network",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("VirtualNetworkPropertiesFormat.EnableVMProtection"),
			},
			{
				Name:        "ddos_protection_plan_id",
				Description: "Resource ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualNetworkPropertiesFormat.DdosProtectionPlan.ID"),
			},
			{
				Name:        "bgp_communities_virtual_network_community",
				Description: "The BGP community associated with the virtual network",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualNetworkPropertiesFormat.BgpCommunities.VirtualNetworkCommunity"),
			},
			{
				Name:        "bgp_communities_regional_community",
				Description: "The BGP community associated with the region of the virtual network",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VirtualNetworkPropertiesFormat.BgpCommunities.RegionalCommunity"),
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "Resource ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_network_virtual_network_subnets",
				Description: "Azure virtual network subnet",
				Resolver:    fetchNetworkVirtualNetworkSubnets,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"virtual_network_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "virtual_network_cq_id",
						Description: "Unique CloudQuery ID of azure_network_virtual_networks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "address_prefix",
						Description: "The address prefix for the subnet",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.AddressPrefix"),
					},
					{
						Name:        "address_prefixes",
						Description: "List of address prefixes for the subnet",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.AddressPrefixes"),
					},
					{
						Name:        "security_group_properties_format_resource_guid",
						Description: "The resource GUID property of the network security group resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.SecurityGroupPropertiesFormat.ResourceGUID"),
					},
					{
						Name:        "security_group_properties_format_provisioning_state",
						Description: "The provisioning state of the network security group resource Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.SecurityGroupPropertiesFormat.ProvisioningState"),
					},
					{
						Name:        "network_security_group_etag",
						Description: "A unique read-only string that changes whenever the resource is updated",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.Etag"),
					},
					{
						Name:        "network_security_group_id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.ID"),
					},
					{
						Name:        "network_security_group_name",
						Description: "Resource name",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.Name"),
					},
					{
						Name:        "network_security_group_type",
						Description: "Resource type",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.Type"),
					},
					{
						Name:        "network_security_group_location",
						Description: "Resource location",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.Location"),
					},
					{
						Name:        "network_security_group_tags",
						Description: "Resource tags",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.NetworkSecurityGroup.Tags"),
					},
					{
						Name:        "route_table_disable_bgp_route_propagation",
						Description: "Whether to disable the routes learned by BGP on that route table True means disable",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.RouteTable.RouteTablePropertiesFormat.DisableBgpRoutePropagation"),
					},
					{
						Name:        "route_table_provisioning_state",
						Description: "The provisioning state of the route table resource Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.RouteTable.RouteTablePropertiesFormat.ProvisioningState"),
					},
					{
						Name:        "route_table_resource_guid",
						Description: "The resource GUID property of the route table",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.RouteTable.RouteTablePropertiesFormat.ResourceGUID"),
					},
					{
						Name:        "route_table_etag",
						Description: "A unique read-only string that changes whenever the resource is updated",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.RouteTable.Etag"),
					},
					{
						Name:        "route_table_id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.RouteTable.ID"),
					},
					{
						Name:        "route_table_name",
						Description: "Resource name",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.RouteTable.Name"),
					},
					{
						Name:        "route_table_type",
						Description: "Resource type",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.RouteTable.Type"),
					},
					{
						Name:        "route_table_location",
						Description: "Resource location",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.RouteTable.Location"),
					},
					{
						Name:        "route_table_tags",
						Description: "Resource tags",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.RouteTable.Tags"),
					},
					{
						Name:        "nat_gateway_id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.NatGateway.ID"),
					},
					{
						Name:        "purpose",
						Description: "A read-only string identifying the intention of use for this subnet based on delegations and other user-defined properties",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.Purpose"),
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the subnet resource Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.ProvisioningState"),
					},
					{
						Name:        "private_endpoint_network_policies",
						Description: "Enable or Disable apply network policies on private end point in the subnet",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.PrivateEndpointNetworkPolicies"),
					},
					{
						Name:        "private_link_service_network_policies",
						Description: "Enable or Disable apply network policies on private link service in the subnet",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetPropertiesFormat.PrivateLinkServiceNetworkPolicies"),
					},
					{
						Name:        "name",
						Description: "The name of the resource that is unique within a resource group This name can be used to access the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "etag",
						Description: "A unique read-only string that changes whenever the resource is updated",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
			},
			{
				Name:        "azure_network_virtual_network_peerings",
				Description: "Azure virtual network peering",
				Resolver:    fetchNetworkVirtualNetworkPeerings,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"virtual_network_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "virtual_network_cq_id",
						Description: "Unique CloudQuery ID of azure_network_virtual_networks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allow_virtual_network_access",
						Description: "Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.AllowVirtualNetworkAccess"),
					},
					{
						Name:        "allow_forwarded_traffic",
						Description: "Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.AllowForwardedTraffic"),
					},
					{
						Name:        "allow_gateway_transit",
						Description: "If gateway links can be used in remote virtual networking to link to this virtual network",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.AllowGatewayTransit"),
					},
					{
						Name:        "use_remote_gateways",
						Description: "If remote gateways can be used on this virtual network If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit Only one peering can have this flag set to true This flag cannot be set if virtual network already has a gateway",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.UseRemoteGateways"),
					},
					{
						Name:        "remote_virtual_network_id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.RemoteVirtualNetwork.ID"),
					},
					{
						Name:        "remote_address_space_address_prefixes",
						Description: "A list of address blocks reserved for this virtual network in CIDR notation",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.RemoteAddressSpace.AddressPrefixes"),
					},
					{
						Name:        "remote_bgp_communities_virtual_network_community",
						Description: "The BGP community associated with the virtual network",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.RemoteBgpCommunities.VirtualNetworkCommunity"),
					},
					{
						Name:        "remote_bgp_communities_regional_community",
						Description: "The BGP community associated with the region of the virtual network",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.RemoteBgpCommunities.RegionalCommunity"),
					},
					{
						Name:        "peering_state",
						Description: "The status of the virtual network peering Possible values include: 'VirtualNetworkPeeringStateInitiated', 'VirtualNetworkPeeringStateConnected', 'VirtualNetworkPeeringStateDisconnected'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.PeeringState"),
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the virtual network peering resource Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkPeeringPropertiesFormat.ProvisioningState"),
					},
					{
						Name:        "name",
						Description: "The name of the resource that is unique within a resource group This name can be used to access the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "etag",
						Description: "A unique read-only string that changes whenever the resource is updated",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
			},
			{
				Name:        "azure_network_virtual_network_ip_allocations",
				Description: "Azure virtual network ip allocation",
				Resolver:    fetchNetworkVirtualNetworkIpAllocations,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"virtual_network_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "virtual_network_cq_id",
						Description: "Unique CloudQuery ID of azure_network_virtual_networks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Resource ID",
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
func fetchNetworkVirtualNetworks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
func fetchNetworkVirtualNetworkSubnets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	vn := parent.Item.(network.VirtualNetwork)
	if vn.Subnets == nil {
		return nil
	}
	res <- *vn.Subnets
	return nil
}
func fetchNetworkVirtualNetworkPeerings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	vn := parent.Item.(network.VirtualNetwork)
	if vn.VirtualNetworkPeerings == nil {
		return nil
	}
	res <- *vn.VirtualNetworkPeerings
	return nil
}
func fetchNetworkVirtualNetworkIpAllocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	vn := parent.Item.(network.VirtualNetwork)
	if vn.IPAllocations == nil {
		return nil
	}
	res <- *vn.IPAllocations
	return nil
}
