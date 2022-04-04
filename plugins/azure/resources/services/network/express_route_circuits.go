package network

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func NetworkExpressRouteCircuits() *schema.Table {
	return &schema.Table{
		Name:          "azure_network_express_route_circuits",
		Description:   "Azure Network Express Route Circuits",
		Resolver:      fetchNetworkExpressRouteCircuits,
		Multiplex:     client.SubscriptionMultiplex,
		DeleteFilter:  client.DeleteSubscriptionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		IgnoreInTests: true,
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
				Name:        "allow_classic_operations",
				Description: "Allow classic operations.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ExpressRouteCircuitPropertiesFormat.AllowClassicOperations"),
			},
			{
				Name:        "bandwidth_in_gbps",
				Description: "The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("ExpressRouteCircuitPropertiesFormat.BandwidthInGbps"),
			},
			{
				Name:        "circuit_provisioning_state",
				Description: "The CircuitProvisioningState state of the resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRouteCircuitPropertiesFormat.CircuitProvisioningState"),
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        schema.TypeString,
			},
			{
				Name:        "express_route_port_id",
				Description: "The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRouteCircuitPropertiesFormat.ExpressRoutePort.ID"),
			},
			{
				Name:        "gateway_manager_etag",
				Description: "The GatewayManager Etag.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRouteCircuitPropertiesFormat.GatewayManagerEtag"),
			},
			{
				Name:        "global_reach_enabled",
				Description: "Flag denoting global reach status.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ExpressRouteCircuitPropertiesFormat.GlobalReachEnabled"),
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
				Description: "The provisioning state of the express route circuit resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRouteCircuitPropertiesFormat.ProvisioningState"),
			},
			{
				Name:        "service_key",
				Description: "The ServiceKey.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRouteCircuitPropertiesFormat.ServiceKey"),
			},
			{
				Name:        "service_provider_notes",
				Description: "The ServiceProviderNotes.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRouteCircuitPropertiesFormat.ServiceProviderNotes"),
			},
			{
				Name:        "service_provider_properties_bandwidth_in_mbps",
				Description: "The BandwidthInMbps.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ExpressRouteCircuitPropertiesFormat.ServiceProviderProperties.BandwidthInMbps"),
			},
			{
				Name:        "service_provider_properties_peering_location",
				Description: "The peering location.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRouteCircuitPropertiesFormat.ServiceProviderProperties.PeeringLocation"),
			},
			{
				Name:        "service_provider_properties_service_provider_name",
				Description: "The serviceProviderName.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRouteCircuitPropertiesFormat.ServiceProviderProperties.ServiceProviderName"),
			},
			{
				Name:        "service_provider_provisioning_state",
				Description: "The ServiceProviderProvisioningState state of the resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExpressRouteCircuitPropertiesFormat.ProvisioningState"),
			},
			{
				Name:        "sku_family",
				Description: "The family of the SKU.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Family"),
			},
			{
				Name:        "sku_name",
				Description: "The name of the SKU.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The tier of the SKU.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:        "stag",
				Description: "The identifier of the circuit traffic. Outer tag for QinQ encapsulation.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ExpressRouteCircuitPropertiesFormat.Stag"),
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
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_network_express_route_circuit_authorizations",
				Description: "Authorization in an ExpressRouteCircuit resource.",
				Resolver:    fetchNetworkExpressRouteCircuitAuthorizations,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"express_route_circuit_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "express_route_circuit_cq_id",
						Description: "Unique CloudQuery ID of azure_network_express_route_circuits table (FK)",
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
						Description: "The authorization key.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthorizationPropertiesFormat.AuthorizationKey"),
					},
					{
						Name:        "authorization_use_status",
						Description: "The authorization use status.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthorizationPropertiesFormat.AuthorizationUseStatus"),
					},
					{
						Name:        "etag",
						Description: "A unique read-only string that changes whenever the resource is updated.",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the authorization resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AuthorizationPropertiesFormat.ProvisioningState"),
					},
					{
						Name:        "type",
						Description: "Resource type.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_network_express_route_circuit_peerings",
				Description: "Peering in an ExpressRouteCircuit resource.",
				Resolver:    fetchNetworkExpressRouteCircuitPeerings,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"express_route_circuit_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "express_route_circuit_cq_id",
						Description: "Unique CloudQuery ID of azure_network_express_route_circuits table (FK)",
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
						Name:        "azure_asn",
						Description: "The Azure ASN.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.AzureASN"),
					},
					{
						Name:        "etag",
						Description: "A unique read-only string that changes whenever the resource is updated.",
						Type:        schema.TypeString,
					},
					{
						Name:        "express_route_connection_id",
						Description: "The ID of the ExpressRouteConnection.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.ExpressRouteConnection.ID"),
					},
					{
						Name:        "gateway_manager_etag",
						Description: "The GatewayManager Etag.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.GatewayManagerEtag"),
					},
					{
						Name:        "ipv6_peering_config_microsoft_peering_config",
						Description: "The Microsoft peering configuration.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.Ipv6PeeringConfig.MicrosoftPeeringConfig"),
					},
					{
						Name:        "ipv6_peering_config_primary_peer_address_prefix",
						Description: "The primary address prefix.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.Ipv6PeeringConfig.PrimaryPeerAddressPrefix"),
					},
					{
						Name:        "ipv6_peering_config_route_filter_id",
						Description: "The reference to the RouteFilter resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.Ipv6PeeringConfig.RouteFilter.ID"),
					},
					{
						Name:        "ipv6_peering_config_secondary_peer_address_prefix",
						Description: "The secondary address prefix.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.Ipv6PeeringConfig.SecondaryPeerAddressPrefix"),
					},
					{
						Name:        "ipv6_peering_config_state",
						Description: "The state of peering.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.Ipv6PeeringConfig.State"),
					},
					{
						Name:        "last_modified_by",
						Description: "Who was the last to modify the peering.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.LastModifiedBy"),
					},
					{
						Name:        "microsoft_peering_config_advertised_communities",
						Description: "The communities of bgp peering. Specified for microsoft peering.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.MicrosoftPeeringConfig.AdvertisedCommunities"),
					},
					{
						Name:        "microsoft_peering_config_advertised_public_prefixes",
						Description: "The reference to AdvertisedPublicPrefixes.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.MicrosoftPeeringConfig.AdvertisedPublicPrefixes"),
					},
					{
						Name:        "microsoft_peering_config_advertised_public_prefixes_state",
						Description: "The advertised public prefix state of the Peering resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.MicrosoftPeeringConfig.AdvertisedPublicPrefixesState"),
					},
					{
						Name:        "microsoft_peering_config_customer_asn",
						Description: "The CustomerASN of the peering.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.MicrosoftPeeringConfig.CustomerASN"),
					},
					{
						Name:        "microsoft_peering_config_legacy_mode",
						Description: "The legacy mode of the peering.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.MicrosoftPeeringConfig.LegacyMode"),
					},
					{
						Name:        "microsoft_peering_config_routing_registry_name",
						Description: "The RoutingRegistryName of the configuration.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.MicrosoftPeeringConfig.RoutingRegistryName"),
					},
					{
						Name:        "name",
						Description: "Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "peer_asn",
						Description: "The peer ASN.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.PeerASN"),
					},
					{
						Name:        "peering_type",
						Description: "The peering type.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.PeeringType"),
					},
					{
						Name:        "primary_azure_port",
						Description: "The primary port.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.PrimaryAzurePort"),
					},
					{
						Name:        "primary_peer_address_prefix",
						Description: "The primary address prefix.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.PrimaryPeerAddressPrefix"),
					},
					{
						Name:        "provisioning_state",
						Description: "The provisioning state of the express route circuit peering resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.ProvisioningState"),
					},
					{
						Name:        "route_filter_id",
						Description: "The reference to the RouteFilter resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.RouteFilter.ID"),
					},
					{
						Name:        "secondary_azure_port",
						Description: "The secondary port.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.SecondaryAzurePort"),
					},
					{
						Name:        "secondary_peer_address_prefix",
						Description: "The secondary address prefix.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.SecondaryPeerAddressPrefix"),
					},
					{
						Name:        "shared_key",
						Description: "The shared key.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.SharedKey"),
					},
					{
						Name:        "state",
						Description: "The peering state.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.State"),
					},
					{
						Name:        "stats_primary_bytes_in",
						Description: "The Primary BytesIn of the peering.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.Stats.PrimarybytesIn"),
					},
					{
						Name:        "stats_primary_bytes_out",
						Description: "The Primary BytesOut of the peering.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.Stats.PrimarybytesOut"),
					},
					{
						Name:        "stats_secondary_bytes_in",
						Description: "The secondary BytesIn of the peering.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.Stats.SecondarybytesIn"),
					},
					{
						Name:        "stats_secondary_bytes_out",
						Description: "The secondary BytesOut of the peering.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.Stats.SecondarybytesOut"),
					},
					{
						Name:        "type",
						Description: "Resource type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "vlan_id",
						Description: "The VLAN ID.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ExpressRouteCircuitPeeringPropertiesFormat.VlanID"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_network_express_route_circuit_connections",
						Description: "Express Route Circuit Connection in an ExpressRouteCircuitPeering resource.",
						Resolver:    fetchNetworkExpressRouteCircuitConnections,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"express_route_circuit_peering_cq_id", "id"}},
						Columns: []schema.Column{
							{
								Name:        "express_route_circuit_peering_cq_id",
								Description: "Unique CloudQuery ID of azure_network_express_route_circuit_peerings table (FK)",
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
								Name:        "address_prefix",
								Description: "/29 IP address space to carve out Customer addresses for tunnels.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ExpressRouteCircuitConnectionPropertiesFormat.AddressPrefix"),
							},
							{
								Name:        "circuit_connection_status",
								Description: "Express Route Circuit connection state.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ExpressRouteCircuitConnectionPropertiesFormat.CircuitConnectionStatus"),
							},
							{
								Name:        "etag",
								Description: "A unique read-only string that changes whenever the resource is updated.",
								Type:        schema.TypeString,
							},
							{
								Name:        "express_route_circuit_peering_id",
								Description: "Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ExpressRouteCircuitConnectionPropertiesFormat.ExpressRouteCircuitPeering.ID"),
							},
							{
								Name:        "ipv6_circuit_connection_config_address_prefix",
								Description: "/125 IP address space to carve out customer addresses for global reach.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ExpressRouteCircuitConnectionPropertiesFormat.Ipv6CircuitConnectionConfig.AddressPrefix"),
							},
							{
								Name:        "ipv6_circuit_connection_config_circuit_connection_status",
								Description: "Express Route Circuit connection state.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ExpressRouteCircuitConnectionPropertiesFormat.Ipv6CircuitConnectionConfig.CircuitConnectionStatus"),
							},
							{
								Name:        "name",
								Description: "Resource name.",
								Type:        schema.TypeString,
							},
							{
								Name:        "peer_express_route_circuit_peering_id",
								Description: "Reference to Express Route Circuit Private Peering Resource of the peered circuit.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ExpressRouteCircuitConnectionPropertiesFormat.PeerExpressRouteCircuitPeering.ID"),
							},
							{
								Name:        "provisioning_state",
								Description: "The provisioning state of the express route circuit connection resource.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ExpressRouteCircuitConnectionPropertiesFormat.ProvisioningState"),
							},
							{
								Name:        "type",
								Description: "Resource type.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "azure_network_peer_express_route_circuit_connections",
						Description: "Peer Express Route Circuit Connection in an ExpressRouteCircuitPeering resource.",
						Resolver:    fetchNetworkPeerExpressRouteCircuitConnections,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"express_route_circuit_peering_cq_id", "id"}},
						Columns: []schema.Column{
							{
								Name:        "express_route_circuit_peering_cq_id",
								Description: "Unique CloudQuery ID of azure_network_express_route_circuit_peerings table (FK)",
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
								Name:        "address_prefix",
								Description: "/29 IP address space to carve out Customer addresses for tunnels.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("PeerExpressRouteCircuitConnectionPropertiesFormat.AddressPrefix"),
							},
							{
								Name:        "auth_resource_guid",
								Description: "The resource guid of the authorization used for the express route circuit connection.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("PeerExpressRouteCircuitConnectionPropertiesFormat.AuthResourceGUID"),
							},
							{
								Name:        "circuit_connection_status",
								Description: "Express Route Circuit connection state.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("PeerExpressRouteCircuitConnectionPropertiesFormat.CircuitConnectionStatus"),
							},
							{
								Name:        "connection_name",
								Description: "The name of the express route circuit connection resource.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("PeerExpressRouteCircuitConnectionPropertiesFormat.ConnectionName"),
							},
							{
								Name:        "etag",
								Description: "A unique read-only string that changes whenever the resource is updated.",
								Type:        schema.TypeString,
							},
							{
								Name:        "express_route_circuit_peering_id",
								Description: "Reference to Express Route Circuit Private Peering Resource of the circuit.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("PeerExpressRouteCircuitConnectionPropertiesFormat.ExpressRouteCircuitPeering.ID"),
							},
							{
								Name:        "name",
								Description: "Resource name.",
								Type:        schema.TypeString,
							},
							{
								Name:        "peer_express_route_circuit_peering_id",
								Description: "Reference to Express Route Circuit Private Peering Resource of the peered circuit.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("PeerExpressRouteCircuitConnectionPropertiesFormat.PeerExpressRouteCircuitPeering.ID"),
							},
							{
								Name:        "provisioning_state",
								Description: "The provisioning state of the peer express route circuit connection resource.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("PeerExpressRouteCircuitConnectionPropertiesFormat.ProvisioningState"),
							},
							{
								Name:        "type",
								Description: "Resource type.",
								Type:        schema.TypeString,
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

func fetchNetworkExpressRouteCircuits(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.ExpressRouteCircuits
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
func fetchNetworkExpressRouteCircuitAuthorizations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	erc, ok := parent.Item.(network.ExpressRouteCircuit)
	if !ok {
		return fmt.Errorf("expected to have network.ExpressRouteCircuit but got %T", parent.Item)
	}
	if erc.ExpressRouteCircuitPropertiesFormat != nil && erc.ExpressRouteCircuitPropertiesFormat.Authorizations != nil {
		res <- *erc.ExpressRouteCircuitPropertiesFormat.Authorizations
	}
	return nil
}
func fetchNetworkExpressRouteCircuitPeerings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	erc, ok := parent.Item.(network.ExpressRouteCircuit)
	if !ok {
		return fmt.Errorf("expected to have network.ExpressRouteCircuit but got %T", parent.Item)
	}
	if erc.ExpressRouteCircuitPropertiesFormat != nil && erc.ExpressRouteCircuitPropertiesFormat.Peerings != nil {
		res <- *erc.ExpressRouteCircuitPropertiesFormat.Peerings
	}
	return nil
}
func fetchNetworkExpressRouteCircuitConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	rcp, ok := parent.Item.(network.ExpressRouteCircuitPeering)
	if !ok {
		return fmt.Errorf("expected to have network.ExpressRouteCircuitPeering but got %T", parent.Item)
	}
	if rcp.ExpressRouteCircuitPeeringPropertiesFormat != nil && rcp.ExpressRouteCircuitPeeringPropertiesFormat.Connections != nil {
		res <- *rcp.ExpressRouteCircuitPeeringPropertiesFormat.Connections
	}
	return nil
}
func fetchNetworkPeerExpressRouteCircuitConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	rcp, ok := parent.Item.(network.ExpressRouteCircuitPeering)
	if !ok {
		return fmt.Errorf("expected to have network.ExpressRouteCircuitPeering but got %T", parent.Item)
	}
	if rcp.ExpressRouteCircuitPeeringPropertiesFormat != nil && rcp.ExpressRouteCircuitPeeringPropertiesFormat.PeeredConnections != nil {
		res <- *rcp.ExpressRouteCircuitPeeringPropertiesFormat.PeeredConnections
	}
	return nil
}
