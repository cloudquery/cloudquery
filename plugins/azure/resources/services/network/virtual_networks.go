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
			{
				Name:        "azure_network_virtual_network_gateways",
				Description: "Azure virtual network gateway",
				Resolver:    fetchNetworksVirtualNetworkGateways,
				Columns: []schema.Column{
					{
						Name:        "virtual_network_cq_id",
						Description: "Unique CloudQuery ID of azure_network_virtual_networks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
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
						Name:        "ip_configurations",
						Description: "IP configurations for virtual network gateway.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.IPConfigurations"),
					},
					{
						Name:        "gateway_type",
						Description: "The type of this virtual network gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.GatewayType"),
					},
					{
						Name:        "vpn_type",
						Description: "The type of this virtual network gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.VpnType"),
					},
					{
						Name:        "vpn_gateway_generation",
						Description: "The generation for this VirtualNetworkGateway.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.VpnGatewayGeneration"),
					},
					{
						Name:        "enable_bgp",
						Description: "Whether BGP is enabled for this virtual network gateway or not.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.EnableBgp"),
					},
					{
						Name:        "enable_private_ip_address",
						Description: "Whether private IP needs to be enabled on this gateway for connections or not.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.EnablePrivateIPAddress"),
					},
					{
						Name:        "active_active",
						Description: "ActiveActive flag.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.ActiveActive"),
					},
					{
						Name:          "gateway_default_site_id",
						Description:   "The reference to the LocalNetworkGateway resource which represents local network site having default routes.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.GatewayDefaultSite.ID"),
						IgnoreInTests: true,
					},
					{
						Name:        "sku_name",
						Description: "Gateway SKU name.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.Sku.Name"),
					},
					{
						Name:        "sku_tier",
						Description: "Gateway SKU tier.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.Sku.Tier"),
					},
					{
						Name:        "sku_capacity",
						Description: "READ-ONLY; The capacity.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.Sku.Capacity"),
					},
					{
						Name:          "vpn_client_configuration_address_pool",
						Description:   "The reference to the address space resource which represents Address space for P2S VpnClient.",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.VpnClientConfiguration.VpnClientAddressPool.AddressPrefixes"),
						IgnoreInTests: true,
					},
					{
						Name:        "vpn_client_configuration_root_certificates",
						Description: "VpnClientRootCertificate for virtual network gateway.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.VpnClientConfiguration.VpnClientRootCertificates"),
					},
					{
						Name:        "vpn_client_configuration_revoked_certificates",
						Description: "VpnClientRevokedCertificate for Virtual network gateway.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.VpnClientConfiguration.VpnClientRevokedCertificates"),
					},
					{
						Name:        "vpn_client_configuration_protocols",
						Description: "VpnClientProtocols for Virtual network gateway.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.VpnClientConfiguration.VpnClientProtocols"),
					},
					{
						Name:        "vpn_client_configuration_authentication_types",
						Description: "VPN authentication types for the virtual network gateway.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.VpnClientConfiguration.VpnAuthenticationTypes"),
					},
					{
						Name:        "vpn_client_configuration_ipsec_policies",
						Description: "VpnClientIpsecPolicies for virtual network gateway P2S client.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.VpnClientConfiguration.VpnClientIpsecPolicies"),
					},
					{
						Name:          "vpn_client_configuration_radius_server_address",
						Description:   "The radius server address property of the VirtualNetworkGateway resource for vpn client connection.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.VpnClientConfiguration.RadiusServerAddress"),
						IgnoreInTests: true,
					},
					{
						Name:          "vpn_client_configuration_radius_server_secret",
						Description:   "The radius secret property of the VirtualNetworkGateway resource for vpn client connection.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.VpnClientConfiguration.RadiusServerSecret"),
						IgnoreInTests: true,
					},
					{
						Name:        "vpn_client_configuration_radius_servers",
						Description: "The radiusServers property for multiple radius server configuration.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.VpnClientConfiguration.RadiusServers"),
					},
					{
						Name:          "vpn_client_configuration_aad_tenant",
						Description:   "The AADTenant property of the VirtualNetworkGateway resource for vpn client connection used for AAD authentication.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.VpnClientConfiguration.AadTenant"),
						IgnoreInTests: true,
					},
					{
						Name:          "vpn_client_configuration_aad_audience",
						Description:   "The AADAudience property of the VirtualNetworkGateway resource for vpn client connection used for AAD authentication.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.VpnClientConfiguration.AadAudience"),
						IgnoreInTests: true,
					},
					{
						Name:          "vpn_client_configuration_aad_issuer",
						Description:   "The AADIssuer property of the VirtualNetworkGateway resource for vpn client connection used for AAD authentication.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.VpnClientConfiguration.AadIssuer"),
						IgnoreInTests: true,
					},
					{
						Name:        "bgp_settings_asn",
						Description: "The BGP speaker's ASN.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.BgpSettings.Asn"),
					},
					{
						Name:        "bgp_settings_bgp_peering_address",
						Description: "The BGP peering address and BGP identifier of this BGP speaker.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.BgpSettings.BgpPeeringAddress"),
					},
					{
						Name:        "bgp_settings_peer_weight",
						Description: "The weight added to routes learned from this BGP speaker.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.BgpSettings.PeerWeight"),
					},
					{
						Name:        "bgp_settings_bgp_peering_addresses",
						Description: "BGP peering address with IP configuration ID for virtual network gateway.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.BgpSettings.BgpPeeringAddresses"),
					},
					{
						Name:          "custom_routes_address_prefixes",
						Description:   "The reference to the address space resource which represents the custom routes address space specified by the customer for virtual network gateway and VpnClient.",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.CustomRoutes.AddressPrefixes"),
						IgnoreInTests: true,
					},
					{
						Name:        "resource_guid",
						Description: "The resource GUID property of the virtual network gateway resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.ResourceGUID"),
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the virtual network gateway resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.ProvisioningState"),
					},
					{
						Name:          "enable_dns_forwarding",
						Description:   "Whether dns forwarding is enabled or not.",
						Type:          schema.TypeBool,
						Resolver:      schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.EnableDNSForwarding"),
						IgnoreInTests: true,
					},
					{
						Name:          "inbound_dns_forwarding_endpoint",
						Description:   "The IP address allocated by the gateway to which dns requests can be sent.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.InboundDNSForwardingEndpoint"),
						IgnoreInTests: true,
					},
					{
						Name:          "vnet_extended_location_resource_id",
						Description:   "Customer vnet resource id.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("VirtualNetworkGatewayPropertiesFormat.VNetExtendedLocationResourceID"),
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
						Name:        "azure_network_virtual_network_gateway_connections",
						Description: "Azure virtual network gateway connection",
						Resolver:    fetchNetworksVirtualNetworkGatewayConnections,
						Columns: []schema.Column{
							{
								Name:        "virtual_network_gateway_cq_id",
								Description: "Unique CloudQuery ID of azure_network_virtual_network_gateways table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:          "authorization_key",
								Description:   "The authorizationKey.",
								Type:          schema.TypeString,
								Resolver:      schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.AuthorizationKey"),
								IgnoreInTests: true,
							},
							{
								Name:        "virtual_network_gateway_1",
								Description: "The reference to virtual network gateway resource.",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.VirtualNetworkGateway1"),
							},
							{
								Name:          "virtual_network_gateway_2",
								Description:   "The reference to virtual network gateway resource.",
								Type:          schema.TypeJSON,
								Resolver:      schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.VirtualNetworkGateway2"),
								IgnoreInTests: true,
							},
							{
								Name:        "local_network_gateway_2",
								Description: "The reference to local network gateway resource.",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.LocalNetworkGateway2"),
							},
							{
								Name:        "connection_type",
								Description: "Gateway connection type.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.ConnectionType"),
							},
							{
								Name:        "connection_protocol",
								Description: "Connection protocol used for this connection.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.ConnectionProtocol"),
							},
							{
								Name:        "routing_weight",
								Description: "The routing weight.",
								Type:        schema.TypeInt,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.RoutingWeight"),
							},
							{
								Name:        "connection_mode",
								Description: "The connection mode for this connection.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.ConnectionMode"),
							},
							{
								Name:          "shared_key",
								Description:   "The IPSec shared key.",
								Type:          schema.TypeString,
								Resolver:      schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.SharedKey"),
								IgnoreInTests: true,
							},
							{
								Name:        "connection_status",
								Description: "Virtual Network Gateway connection status.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.ConnectionStatus"),
							},
							{
								Name:          "tunnel_connection_status",
								Description:   "Collection of all tunnels' connection health status.",
								Type:          schema.TypeJSON,
								Resolver:      schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.TunnelConnectionStatus"),
								IgnoreInTests: true,
							},
							{
								Name:        "egress_bytes_transferred",
								Description: "The egress bytes transferred in this connection.",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.EgressBytesTransferred"),
							},
							{
								Name:        "ingress_bytes_transferred",
								Description: "The ingress bytes transferred in this connection.",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.IngressBytesTransferred"),
							},
							{
								Name:          "peer_id",
								Description:   "The reference to peerings resource.",
								Type:          schema.TypeString,
								Resolver:      schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.Peer.ID"),
								IgnoreInTests: true,
							},
							{
								Name:        "enable_bgp",
								Description: "EnableBgp flag.",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.EnableBgp"),
							},
							{
								Name:        "use_policy_based_traffic_selectors",
								Description: "Enable policy-based traffic selectors.",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.UsePolicyBasedTrafficSelectors"),
							},
							{
								Name:        "ipsec_policies",
								Description: "The IPSec Policies to be considered by this connection.",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.IpsecPolicies"),
							},
							{
								Name:        "traffic_selector_policies",
								Description: "The Traffic Selector Policies to be considered by this connection.",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.TrafficSelectorPolicies"),
							},
							{
								Name:        "resource_guid",
								Description: "The resource GUID property of the virtual network gateway connection resource.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.ResourceGUID"),
							},
							{
								Name:        "provisioning_state",
								Description: "The provisioning state of the virtual network gateway connection resource.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.ProvisioningState"),
							},
							{
								Name:        "express_route_gateway_bypass",
								Description: "Bypass ExpressRoute Gateway for data forwarding.",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("VirtualNetworkGatewayConnectionListEntityPropertiesFormat.ExpressRouteGatewayBypass"),
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
	vn := resource.Item.(network.VirtualNetwork)
	if vn.DhcpOptions == nil || vn.DhcpOptions.DNSServers == nil || len(*vn.DhcpOptions.DNSServers) == 0 {
		return nil
	}
	ips := make([]net.IP, 0, len(*vn.DhcpOptions.DNSServers))
	for _, ip := range *vn.DhcpOptions.DNSServers {
		i := net.ParseIP(ip)
		if i == nil {
			return diag.WrapError(fmt.Errorf("wrong format of IP: %s", ip))
		}
		ips = append(ips, i)
		net.ParseIP(ip)
	}
	return diag.WrapError(resource.Set(c.Name, ips))
}
func resolveNetworkVirtualNetworksIpAllocations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	vn := resource.Item.(network.VirtualNetwork)
	if vn.IPAllocations == nil {
		return nil
	}
	allocations := make([]string, 0, len(*vn.IPAllocations))
	for _, a := range *vn.IPAllocations {
		allocations = append(allocations, *a.ID)
	}
	return diag.WrapError(resource.Set(c.Name, allocations))
}
func fetchNetworksVirtualNetworkSubnets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	vn := parent.Item.(network.VirtualNetwork)
	if vn.Subnets == nil {
		return nil
	}
	res <- *vn.Subnets
	return nil
}

func fetchNetworksVirtualNetworkPeerings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	vn := parent.Item.(network.VirtualNetwork)
	if vn.VirtualNetworkPeerings == nil {
		return nil
	}
	res <- *vn.VirtualNetworkPeerings
	return nil
}

func fetchNetworksVirtualNetworkGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.VirtualNetworkGateways
	vn := parent.Item.(network.VirtualNetwork)
	details, err := client.ParseResourceID(*vn.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	response, err := svc.List(ctx, details.ResourceGroup)
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

func fetchNetworksVirtualNetworkGatewayConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.VirtualNetworkGateways
	vng := parent.Item.(network.VirtualNetworkGateway)
	details, err := client.ParseResourceID(*vng.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	response, err := svc.ListConnections(ctx, details.ResourceGroup, *vng.Name)
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
