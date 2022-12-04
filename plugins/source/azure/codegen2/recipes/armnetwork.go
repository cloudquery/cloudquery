// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"

func Armnetwork() []Table {
	tables := []Table{
		{
      Name: "endpoint_service_result",
      Struct: &armnetwork.EndpointServiceResult{},
      ResponseStruct: &armnetwork.AvailableEndpointServicesClientListResponse{},
      Client: &armnetwork.AvailableEndpointServicesClient{},
      ListFunc: (&armnetwork.AvailableEndpointServicesClient{}).NewListPager,
			NewFunc: armnetwork.NewAvailableEndpointServicesClient,
		},
		{
      Name: "bgp_service_community",
      Struct: &armnetwork.BgpServiceCommunity{},
      ResponseStruct: &armnetwork.BgpServiceCommunitiesClientListResponse{},
      Client: &armnetwork.BgpServiceCommunitiesClient{},
      ListFunc: (&armnetwork.BgpServiceCommunitiesClient{}).NewListPager,
			NewFunc: armnetwork.NewBgpServiceCommunitiesClient,
		},
		{
      Name: "hub_virtual_network_connection",
      Struct: &armnetwork.HubVirtualNetworkConnection{},
      ResponseStruct: &armnetwork.HubVirtualNetworkConnectionsClientListResponse{},
      Client: &armnetwork.HubVirtualNetworkConnectionsClient{},
      ListFunc: (&armnetwork.HubVirtualNetworkConnectionsClient{}).NewListPager,
			NewFunc: armnetwork.NewHubVirtualNetworkConnectionsClient,
		},
		{
      Name: "frontend_ip_configuration",
      Struct: &armnetwork.FrontendIPConfiguration{},
      ResponseStruct: &armnetwork.LoadBalancerFrontendIPConfigurationsClientListResponse{},
      Client: &armnetwork.LoadBalancerFrontendIPConfigurationsClient{},
      ListFunc: (&armnetwork.LoadBalancerFrontendIPConfigurationsClient{}).NewListPager,
			NewFunc: armnetwork.NewLoadBalancerFrontendIPConfigurationsClient,
		},
		{
      Name: "subnet",
      Struct: &armnetwork.Subnet{},
      ResponseStruct: &armnetwork.SubnetsClientListResponse{},
      Client: &armnetwork.SubnetsClient{},
      ListFunc: (&armnetwork.SubnetsClient{}).NewListPager,
			NewFunc: armnetwork.NewSubnetsClient,
		},
		{
      Name: "custom_ip_prefix",
      Struct: &armnetwork.CustomIPPrefix{},
      ResponseStruct: &armnetwork.CustomIPPrefixesClientListResponse{},
      Client: &armnetwork.CustomIPPrefixesClient{},
      ListFunc: (&armnetwork.CustomIPPrefixesClient{}).NewListPager,
			NewFunc: armnetwork.NewCustomIPPrefixesClient,
		},
		{
      Name: "manager_connection",
      Struct: &armnetwork.ManagerConnection{},
      ResponseStruct: &armnetwork.ManagementGroupNetworkManagerConnectionsClientListResponse{},
      Client: &armnetwork.ManagementGroupNetworkManagerConnectionsClient{},
      ListFunc: (&armnetwork.ManagementGroupNetworkManagerConnectionsClient{}).NewListPager,
			NewFunc: armnetwork.NewManagementGroupNetworkManagerConnectionsClient,
		},
		{
      Name: "available_service_alias",
      Struct: &armnetwork.AvailableServiceAlias{},
      ResponseStruct: &armnetwork.AvailableServiceAliasesClientListResponse{},
      Client: &armnetwork.AvailableServiceAliasesClient{},
      ListFunc: (&armnetwork.AvailableServiceAliasesClient{}).NewListPager,
			NewFunc: armnetwork.NewAvailableServiceAliasesClient,
		},
		{
      Name: "interface_ip_configuration",
      Struct: &armnetwork.InterfaceIPConfiguration{},
      ResponseStruct: &armnetwork.InterfaceIPConfigurationsClientListResponse{},
      Client: &armnetwork.InterfaceIPConfigurationsClient{},
      ListFunc: (&armnetwork.InterfaceIPConfigurationsClient{}).NewListPager,
			NewFunc: armnetwork.NewInterfaceIPConfigurationsClient,
		},
		{
      Name: "local_network_gateway",
      Struct: &armnetwork.LocalNetworkGateway{},
      ResponseStruct: &armnetwork.LocalNetworkGatewaysClientListResponse{},
      Client: &armnetwork.LocalNetworkGatewaysClient{},
      ListFunc: (&armnetwork.LocalNetworkGatewaysClient{}).NewListPager,
			NewFunc: armnetwork.NewLocalNetworkGatewaysClient,
		},
		{
      Name: "virtual_appliance_sku",
      Struct: &armnetwork.VirtualApplianceSKU{},
      ResponseStruct: &armnetwork.VirtualApplianceSKUsClientListResponse{},
      Client: &armnetwork.VirtualApplianceSKUsClient{},
      ListFunc: (&armnetwork.VirtualApplianceSKUsClient{}).NewListPager,
			NewFunc: armnetwork.NewVirtualApplianceSKUsClient,
		},
		{
      Name: "watcher",
      Struct: &armnetwork.Watcher{},
      ResponseStruct: &armnetwork.WatchersClientListResponse{},
      Client: &armnetwork.WatchersClient{},
      ListFunc: (&armnetwork.WatchersClient{}).NewListPager,
			NewFunc: armnetwork.NewWatchersClient,
		},
		{
      Name: "public_ip_address",
      Struct: &armnetwork.PublicIPAddress{},
      ResponseStruct: &armnetwork.PublicIPAddressesClientListResponse{},
      Client: &armnetwork.PublicIPAddressesClient{},
      ListFunc: (&armnetwork.PublicIPAddressesClient{}).NewListPager,
			NewFunc: armnetwork.NewPublicIPAddressesClient,
		},
		{
      Name: "virtual_appliance_site",
      Struct: &armnetwork.VirtualApplianceSite{},
      ResponseStruct: &armnetwork.VirtualApplianceSitesClientListResponse{},
      Client: &armnetwork.VirtualApplianceSitesClient{},
      ListFunc: (&armnetwork.VirtualApplianceSitesClient{}).NewListPager,
			NewFunc: armnetwork.NewVirtualApplianceSitesClient,
		},
		{
      Name: "bgp_connection",
      Struct: &armnetwork.BgpConnection{},
      ResponseStruct: &armnetwork.VirtualHubBgpConnectionsClientListResponse{},
      Client: &armnetwork.VirtualHubBgpConnectionsClient{},
      ListFunc: (&armnetwork.VirtualHubBgpConnectionsClient{}).NewListPager,
			NewFunc: armnetwork.NewVirtualHubBgpConnectionsClient,
		},
		{
      Name: "p_2_svpn_gateway",
      Struct: &armnetwork.P2SVPNGateway{},
      ResponseStruct: &armnetwork.P2SVPNGatewaysClientListResponse{},
      Client: &armnetwork.P2SVPNGatewaysClient{},
      ListFunc: (&armnetwork.P2SVPNGatewaysClient{}).NewListPager,
			NewFunc: armnetwork.NewP2SVPNGatewaysClient,
		},
		{
      Name: "peer_express_route_circuit_connection",
      Struct: &armnetwork.PeerExpressRouteCircuitConnection{},
      ResponseStruct: &armnetwork.PeerExpressRouteCircuitConnectionsClientListResponse{},
      Client: &armnetwork.PeerExpressRouteCircuitConnectionsClient{},
      ListFunc: (&armnetwork.PeerExpressRouteCircuitConnectionsClient{}).NewListPager,
			NewFunc: armnetwork.NewPeerExpressRouteCircuitConnectionsClient,
		},
		{
      Name: "service_endpoint_policy",
      Struct: &armnetwork.ServiceEndpointPolicy{},
      ResponseStruct: &armnetwork.ServiceEndpointPoliciesClientListResponse{},
      Client: &armnetwork.ServiceEndpointPoliciesClient{},
      ListFunc: (&armnetwork.ServiceEndpointPoliciesClient{}).NewListPager,
			NewFunc: armnetwork.NewServiceEndpointPoliciesClient,
		},
		{
      Name: "static_member",
      Struct: &armnetwork.StaticMember{},
      ResponseStruct: &armnetwork.StaticMembersClientListResponse{},
      Client: &armnetwork.StaticMembersClient{},
      ListFunc: (&armnetwork.StaticMembersClient{}).NewListPager,
			NewFunc: armnetwork.NewStaticMembersClient,
		},
		{
      Name: "virtual_hub_route_table_v_2",
      Struct: &armnetwork.VirtualHubRouteTableV2{},
      ResponseStruct: &armnetwork.VirtualHubRouteTableV2SClientListResponse{},
      Client: &armnetwork.VirtualHubRouteTableV2SClient{},
      ListFunc: (&armnetwork.VirtualHubRouteTableV2SClient{}).NewListPager,
			NewFunc: armnetwork.NewVirtualHubRouteTableV2SClient,
		},
		{
      Name: "virtual_wan",
      Struct: &armnetwork.VirtualWAN{},
      ResponseStruct: &armnetwork.VirtualWansClientListResponse{},
      Client: &armnetwork.VirtualWansClient{},
      ListFunc: (&armnetwork.VirtualWansClient{}).NewListPager,
			NewFunc: armnetwork.NewVirtualWansClient,
		},
		{
      Name: "service_tag_information",
      Struct: &armnetwork.ServiceTagInformation{},
      ResponseStruct: &armnetwork.ServiceTagInformationClientListResponse{},
      Client: &armnetwork.ServiceTagInformationClient{},
      ListFunc: (&armnetwork.ServiceTagInformationClient{}).NewListPager,
			NewFunc: armnetwork.NewServiceTagInformationClient,
		},
		{
      Name: "express_route_circuit_connection",
      Struct: &armnetwork.ExpressRouteCircuitConnection{},
      ResponseStruct: &armnetwork.ExpressRouteCircuitConnectionsClientListResponse{},
      Client: &armnetwork.ExpressRouteCircuitConnectionsClient{},
      ListFunc: (&armnetwork.ExpressRouteCircuitConnectionsClient{}).NewListPager,
			NewFunc: armnetwork.NewExpressRouteCircuitConnectionsClient,
		},
		{
      Name: "outbound_rule",
      Struct: &armnetwork.OutboundRule{},
      ResponseStruct: &armnetwork.LoadBalancerOutboundRulesClientListResponse{},
      Client: &armnetwork.LoadBalancerOutboundRulesClient{},
      ListFunc: (&armnetwork.LoadBalancerOutboundRulesClient{}).NewListPager,
			NewFunc: armnetwork.NewLoadBalancerOutboundRulesClient,
		},
		{
      Name: "interface_tap_configuration",
      Struct: &armnetwork.InterfaceTapConfiguration{},
      ResponseStruct: &armnetwork.InterfaceTapConfigurationsClientListResponse{},
      Client: &armnetwork.InterfaceTapConfigurationsClient{},
      ListFunc: (&armnetwork.InterfaceTapConfigurationsClient{}).NewListPager,
			NewFunc: armnetwork.NewInterfaceTapConfigurationsClient,
		},
		{
      Name: "route",
      Struct: &armnetwork.Route{},
      ResponseStruct: &armnetwork.RoutesClientListResponse{},
      Client: &armnetwork.RoutesClient{},
      ListFunc: (&armnetwork.RoutesClient{}).NewListPager,
			NewFunc: armnetwork.NewRoutesClient,
		},
		{
      Name: "available_delegation",
      Struct: &armnetwork.AvailableDelegation{},
      ResponseStruct: &armnetwork.AvailableDelegationsClientListResponse{},
      Client: &armnetwork.AvailableDelegationsClient{},
      ListFunc: (&armnetwork.AvailableDelegationsClient{}).NewListPager,
			NewFunc: armnetwork.NewAvailableDelegationsClient,
		},
		{
      Name: "azure_firewall",
      Struct: &armnetwork.AzureFirewall{},
      ResponseStruct: &armnetwork.AzureFirewallsClientListResponse{},
      Client: &armnetwork.AzureFirewallsClient{},
      ListFunc: (&armnetwork.AzureFirewallsClient{}).NewListPager,
			NewFunc: armnetwork.NewAzureFirewallsClient,
		},
		{
      Name: "packet_capture_result",
      Struct: &armnetwork.PacketCaptureResult{},
      ResponseStruct: &armnetwork.PacketCapturesClientListResponse{},
      Client: &armnetwork.PacketCapturesClient{},
      ListFunc: (&armnetwork.PacketCapturesClient{}).NewListPager,
			NewFunc: armnetwork.NewPacketCapturesClient,
		},
		{
      Name: "route_table",
      Struct: &armnetwork.RouteTable{},
      ResponseStruct: &armnetwork.RouteTablesClientListResponse{},
      Client: &armnetwork.RouteTablesClient{},
      ListFunc: (&armnetwork.RouteTablesClient{}).NewListPager,
			NewFunc: armnetwork.NewRouteTablesClient,
		},
		{
      Name: "security_partner_provider",
      Struct: &armnetwork.SecurityPartnerProvider{},
      ResponseStruct: &armnetwork.SecurityPartnerProvidersClientListResponse{},
      Client: &armnetwork.SecurityPartnerProvidersClient{},
      ListFunc: (&armnetwork.SecurityPartnerProvidersClient{}).NewListPager,
			NewFunc: armnetwork.NewSecurityPartnerProvidersClient,
		},
		{
      Name: "virtual_network_gateway_connection",
      Struct: &armnetwork.VirtualNetworkGatewayConnection{},
      ResponseStruct: &armnetwork.VirtualNetworkGatewayConnectionsClientListResponse{},
      Client: &armnetwork.VirtualNetworkGatewayConnectionsClient{},
      ListFunc: (&armnetwork.VirtualNetworkGatewayConnectionsClient{}).NewListPager,
			NewFunc: armnetwork.NewVirtualNetworkGatewayConnectionsClient,
		},
		{
      Name: "virtual_network_peering",
      Struct: &armnetwork.VirtualNetworkPeering{},
      ResponseStruct: &armnetwork.VirtualNetworkPeeringsClientListResponse{},
      Client: &armnetwork.VirtualNetworkPeeringsClient{},
      ListFunc: (&armnetwork.VirtualNetworkPeeringsClient{}).NewListPager,
			NewFunc: armnetwork.NewVirtualNetworkPeeringsClient,
		},
		{
      Name: "security_rule",
      Struct: &armnetwork.SecurityRule{},
      ResponseStruct: &armnetwork.DefaultSecurityRulesClientListResponse{},
      Client: &armnetwork.DefaultSecurityRulesClient{},
      ListFunc: (&armnetwork.DefaultSecurityRulesClient{}).NewListPager,
			NewFunc: armnetwork.NewDefaultSecurityRulesClient,
		},
		{
      Name: "manager",
      Struct: &armnetwork.Manager{},
      ResponseStruct: &armnetwork.ManagersClientListResponse{},
      Client: &armnetwork.ManagersClient{},
      ListFunc: (&armnetwork.ManagersClient{}).NewListPager,
			NewFunc: armnetwork.NewManagersClient,
		},
		{
      Name: "bastion_host",
      Struct: &armnetwork.BastionHost{},
      ResponseStruct: &armnetwork.BastionHostsClientListResponse{},
      Client: &armnetwork.BastionHostsClient{},
      ListFunc: (&armnetwork.BastionHostsClient{}).NewListPager,
			NewFunc: armnetwork.NewBastionHostsClient,
		},
		{
      Name: "ddos_protection_plan",
      Struct: &armnetwork.DdosProtectionPlan{},
      ResponseStruct: &armnetwork.DdosProtectionPlansClientListResponse{},
      Client: &armnetwork.DdosProtectionPlansClient{},
      ListFunc: (&armnetwork.DdosProtectionPlansClient{}).NewListPager,
			NewFunc: armnetwork.NewDdosProtectionPlansClient,
		},
		{
      Name: "firewall_policy",
      Struct: &armnetwork.FirewallPolicy{},
      ResponseStruct: &armnetwork.FirewallPoliciesClientListResponse{},
      Client: &armnetwork.FirewallPoliciesClient{},
      ListFunc: (&armnetwork.FirewallPoliciesClient{}).NewListPager,
			NewFunc: armnetwork.NewFirewallPoliciesClient,
		},
		{
      Name: "virtual_hub",
      Struct: &armnetwork.VirtualHub{},
      ResponseStruct: &armnetwork.VirtualHubsClientListResponse{},
      Client: &armnetwork.VirtualHubsClient{},
      ListFunc: (&armnetwork.VirtualHubsClient{}).NewListPager,
			NewFunc: armnetwork.NewVirtualHubsClient,
		},
		{
      Name: "virtual_network_gateway",
      Struct: &armnetwork.VirtualNetworkGateway{},
      ResponseStruct: &armnetwork.VirtualNetworkGatewaysClientListResponse{},
      Client: &armnetwork.VirtualNetworkGatewaysClient{},
      ListFunc: (&armnetwork.VirtualNetworkGatewaysClient{}).NewListPager,
			NewFunc: armnetwork.NewVirtualNetworkGatewaysClient,
		},
		{
      Name: "express_route_cross_connection_peering",
      Struct: &armnetwork.ExpressRouteCrossConnectionPeering{},
      ResponseStruct: &armnetwork.ExpressRouteCrossConnectionPeeringsClientListResponse{},
      Client: &armnetwork.ExpressRouteCrossConnectionPeeringsClient{},
      ListFunc: (&armnetwork.ExpressRouteCrossConnectionPeeringsClient{}).NewListPager,
			NewFunc: armnetwork.NewExpressRouteCrossConnectionPeeringsClient,
		},
		{
      Name: "load_balancer",
      Struct: &armnetwork.LoadBalancer{},
      ResponseStruct: &armnetwork.InterfaceLoadBalancersClientListResponse{},
      Client: &armnetwork.InterfaceLoadBalancersClient{},
      ListFunc: (&armnetwork.InterfaceLoadBalancersClient{}).NewListPager,
			NewFunc: armnetwork.NewInterfaceLoadBalancersClient,
		},
		{
      Name: "manager_connection",
      Struct: &armnetwork.ManagerConnection{},
      ResponseStruct: &armnetwork.SubscriptionNetworkManagerConnectionsClientListResponse{},
      Client: &armnetwork.SubscriptionNetworkManagerConnectionsClient{},
      ListFunc: (&armnetwork.SubscriptionNetworkManagerConnectionsClient{}).NewListPager,
			NewFunc: armnetwork.NewSubscriptionNetworkManagerConnectionsClient,
		},
		{
      Name: "connection_monitor_result",
      Struct: &armnetwork.ConnectionMonitorResult{},
      ResponseStruct: &armnetwork.ConnectionMonitorsClientListResponse{},
      Client: &armnetwork.ConnectionMonitorsClient{},
      ListFunc: (&armnetwork.ConnectionMonitorsClient{}).NewListPager,
			NewFunc: armnetwork.NewConnectionMonitorsClient,
		},
		{
      Name: "dscp_configuration",
      Struct: &armnetwork.DscpConfiguration{},
      ResponseStruct: &armnetwork.DscpConfigurationClientListResponse{},
      Client: &armnetwork.DscpConfigurationClient{},
      ListFunc: (&armnetwork.DscpConfigurationClient{}).NewListPager,
			NewFunc: armnetwork.NewDscpConfigurationClient,
		},
		{
      Name: "express_route_ports_location",
      Struct: &armnetwork.ExpressRoutePortsLocation{},
      ResponseStruct: &armnetwork.ExpressRoutePortsLocationsClientListResponse{},
      Client: &armnetwork.ExpressRoutePortsLocationsClient{},
      ListFunc: (&armnetwork.ExpressRoutePortsLocationsClient{}).NewListPager,
			NewFunc: armnetwork.NewExpressRoutePortsLocationsClient,
		},
		{
      Name: "private_endpoint",
      Struct: &armnetwork.PrivateEndpoint{},
      ResponseStruct: &armnetwork.PrivateEndpointsClientListResponse{},
      Client: &armnetwork.PrivateEndpointsClient{},
      ListFunc: (&armnetwork.PrivateEndpointsClient{}).NewListPager,
			NewFunc: armnetwork.NewPrivateEndpointsClient,
		},
		{
      Name: "usage",
      Struct: &armnetwork.Usage{},
      ResponseStruct: &armnetwork.UsagesClientListResponse{},
      Client: &armnetwork.UsagesClient{},
      ListFunc: (&armnetwork.UsagesClient{}).NewListPager,
			NewFunc: armnetwork.NewUsagesClient,
		},
		{
      Name: "vpn_server_configuration",
      Struct: &armnetwork.VPNServerConfiguration{},
      ResponseStruct: &armnetwork.VPNServerConfigurationsClientListResponse{},
      Client: &armnetwork.VPNServerConfigurationsClient{},
      ListFunc: (&armnetwork.VPNServerConfigurationsClient{}).NewListPager,
			NewFunc: armnetwork.NewVPNServerConfigurationsClient,
		},
		{
      Name: "available_private_endpoint_type",
      Struct: &armnetwork.AvailablePrivateEndpointType{},
      ResponseStruct: &armnetwork.AvailablePrivateEndpointTypesClientListResponse{},
      Client: &armnetwork.AvailablePrivateEndpointTypesClient{},
      ListFunc: (&armnetwork.AvailablePrivateEndpointTypesClient{}).NewListPager,
			NewFunc: armnetwork.NewAvailablePrivateEndpointTypesClient,
		},
		{
      Name: "available_delegation",
      Struct: &armnetwork.AvailableDelegation{},
      ResponseStruct: &armnetwork.AvailableResourceGroupDelegationsClientListResponse{},
      Client: &armnetwork.AvailableResourceGroupDelegationsClient{},
      ListFunc: (&armnetwork.AvailableResourceGroupDelegationsClient{}).NewListPager,
			NewFunc: armnetwork.NewAvailableResourceGroupDelegationsClient,
		},
		{
      Name: "hub_route_table",
      Struct: &armnetwork.HubRouteTable{},
      ResponseStruct: &armnetwork.HubRouteTablesClientListResponse{},
      Client: &armnetwork.HubRouteTablesClient{},
      ListFunc: (&armnetwork.HubRouteTablesClient{}).NewListPager,
			NewFunc: armnetwork.NewHubRouteTablesClient,
		},
		{
      Name: "ip_group",
      Struct: &armnetwork.IPGroup{},
      ResponseStruct: &armnetwork.IPGroupsClientListResponse{},
      Client: &armnetwork.IPGroupsClient{},
      ListFunc: (&armnetwork.IPGroupsClient{}).NewListPager,
			NewFunc: armnetwork.NewIPGroupsClient,
		},
		{
      Name: "operation",
      Struct: &armnetwork.Operation{},
      ResponseStruct: &armnetwork.OperationsClientListResponse{},
      Client: &armnetwork.OperationsClient{},
      ListFunc: (&armnetwork.OperationsClient{}).NewListPager,
			NewFunc: armnetwork.NewOperationsClient,
		},
		{
      Name: "security_admin_configuration",
      Struct: &armnetwork.SecurityAdminConfiguration{},
      ResponseStruct: &armnetwork.SecurityAdminConfigurationsClientListResponse{},
      Client: &armnetwork.SecurityAdminConfigurationsClient{},
      ListFunc: (&armnetwork.SecurityAdminConfigurationsClient{}).NewListPager,
			NewFunc: armnetwork.NewSecurityAdminConfigurationsClient,
		},
		{
      Name: "express_route_circuit",
      Struct: &armnetwork.ExpressRouteCircuit{},
      ResponseStruct: &armnetwork.ExpressRouteCircuitsClientListResponse{},
      Client: &armnetwork.ExpressRouteCircuitsClient{},
      ListFunc: (&armnetwork.ExpressRouteCircuitsClient{}).NewListPager,
			NewFunc: armnetwork.NewExpressRouteCircuitsClient,
		},
		{
      Name: "backend_address_pool",
      Struct: &armnetwork.BackendAddressPool{},
      ResponseStruct: &armnetwork.LoadBalancerBackendAddressPoolsClientListResponse{},
      Client: &armnetwork.LoadBalancerBackendAddressPoolsClient{},
      ListFunc: (&armnetwork.LoadBalancerBackendAddressPoolsClient{}).NewListPager,
			NewFunc: armnetwork.NewLoadBalancerBackendAddressPoolsClient,
		},
		{
      Name: "private_dns_zone_group",
      Struct: &armnetwork.PrivateDNSZoneGroup{},
      ResponseStruct: &armnetwork.PrivateDNSZoneGroupsClientListResponse{},
      Client: &armnetwork.PrivateDNSZoneGroupsClient{},
      ListFunc: (&armnetwork.PrivateDNSZoneGroupsClient{}).NewListPager,
			NewFunc: armnetwork.NewPrivateDNSZoneGroupsClient,
		},
		{
      Name: "express_route_circuit_peering",
      Struct: &armnetwork.ExpressRouteCircuitPeering{},
      ResponseStruct: &armnetwork.ExpressRouteCircuitPeeringsClientListResponse{},
      Client: &armnetwork.ExpressRouteCircuitPeeringsClient{},
      ListFunc: (&armnetwork.ExpressRouteCircuitPeeringsClient{}).NewListPager,
			NewFunc: armnetwork.NewExpressRouteCircuitPeeringsClient,
		},
		{
      Name: "virtual_appliance",
      Struct: &armnetwork.VirtualAppliance{},
      ResponseStruct: &armnetwork.VirtualAppliancesClientListResponse{},
      Client: &armnetwork.VirtualAppliancesClient{},
      ListFunc: (&armnetwork.VirtualAppliancesClient{}).NewListPager,
			NewFunc: armnetwork.NewVirtualAppliancesClient,
		},
		{
      Name: "vpn_site",
      Struct: &armnetwork.VPNSite{},
      ResponseStruct: &armnetwork.VPNSitesClientListResponse{},
      Client: &armnetwork.VPNSitesClient{},
      ListFunc: (&armnetwork.VPNSitesClient{}).NewListPager,
			NewFunc: armnetwork.NewVPNSitesClient,
		},
		{
      Name: "load_balancing_rule",
      Struct: &armnetwork.LoadBalancingRule{},
      ResponseStruct: &armnetwork.LoadBalancerLoadBalancingRulesClientListResponse{},
      Client: &armnetwork.LoadBalancerLoadBalancingRulesClient{},
      ListFunc: (&armnetwork.LoadBalancerLoadBalancingRulesClient{}).NewListPager,
			NewFunc: armnetwork.NewLoadBalancerLoadBalancingRulesClient,
		},
		{
      Name: "nat_gateway",
      Struct: &armnetwork.NatGateway{},
      ResponseStruct: &armnetwork.NatGatewaysClientListResponse{},
      Client: &armnetwork.NatGatewaysClient{},
      ListFunc: (&armnetwork.NatGatewaysClient{}).NewListPager,
			NewFunc: armnetwork.NewNatGatewaysClient,
		},
		{
      Name: "virtual_router_peering",
      Struct: &armnetwork.VirtualRouterPeering{},
      ResponseStruct: &armnetwork.VirtualRouterPeeringsClientListResponse{},
      Client: &armnetwork.VirtualRouterPeeringsClient{},
      ListFunc: (&armnetwork.VirtualRouterPeeringsClient{}).NewListPager,
			NewFunc: armnetwork.NewVirtualRouterPeeringsClient,
		},
		{
      Name: "application_gateway_private_link_resource",
      Struct: &armnetwork.ApplicationGatewayPrivateLinkResource{},
      ResponseStruct: &armnetwork.ApplicationGatewayPrivateLinkResourcesClientListResponse{},
      Client: &armnetwork.ApplicationGatewayPrivateLinkResourcesClient{},
      ListFunc: (&armnetwork.ApplicationGatewayPrivateLinkResourcesClient{}).NewListPager,
			NewFunc: armnetwork.NewApplicationGatewayPrivateLinkResourcesClient,
		},
		{
      Name: "express_route_link",
      Struct: &armnetwork.ExpressRouteLink{},
      ResponseStruct: &armnetwork.ExpressRouteLinksClientListResponse{},
      Client: &armnetwork.ExpressRouteLinksClient{},
      ListFunc: (&armnetwork.ExpressRouteLinksClient{}).NewListPager,
			NewFunc: armnetwork.NewExpressRouteLinksClient,
		},
		{
      Name: "routing_intent",
      Struct: &armnetwork.RoutingIntent{},
      ResponseStruct: &armnetwork.RoutingIntentClientListResponse{},
      Client: &armnetwork.RoutingIntentClient{},
      ListFunc: (&armnetwork.RoutingIntentClient{}).NewListPager,
			NewFunc: armnetwork.NewRoutingIntentClient,
		},
		{
      Name: "hub_ip_configuration",
      Struct: &armnetwork.HubIPConfiguration{},
      ResponseStruct: &armnetwork.VirtualHubIPConfigurationClientListResponse{},
      Client: &armnetwork.VirtualHubIPConfigurationClient{},
      ListFunc: (&armnetwork.VirtualHubIPConfigurationClient{}).NewListPager,
			NewFunc: armnetwork.NewVirtualHubIPConfigurationClient,
		},
		{
      Name: "flow_log",
      Struct: &armnetwork.FlowLog{},
      ResponseStruct: &armnetwork.FlowLogsClientListResponse{},
      Client: &armnetwork.FlowLogsClient{},
      ListFunc: (&armnetwork.FlowLogsClient{}).NewListPager,
			NewFunc: armnetwork.NewFlowLogsClient,
		},
		{
      Name: "inbound_nat_rule",
      Struct: &armnetwork.InboundNatRule{},
      ResponseStruct: &armnetwork.InboundNatRulesClientListResponse{},
      Client: &armnetwork.InboundNatRulesClient{},
      ListFunc: (&armnetwork.InboundNatRulesClient{}).NewListPager,
			NewFunc: armnetwork.NewInboundNatRulesClient,
		},
		{
      Name: "load_balancer",
      Struct: &armnetwork.LoadBalancer{},
      ResponseStruct: &armnetwork.LoadBalancersClientListResponse{},
      Client: &armnetwork.LoadBalancersClient{},
      ListFunc: (&armnetwork.LoadBalancersClient{}).NewListPager,
			NewFunc: armnetwork.NewLoadBalancersClient,
		},
		{
      Name: "scope_connection",
      Struct: &armnetwork.ScopeConnection{},
      ResponseStruct: &armnetwork.ScopeConnectionsClientListResponse{},
      Client: &armnetwork.ScopeConnectionsClient{},
      ListFunc: (&armnetwork.ScopeConnectionsClient{}).NewListPager,
			NewFunc: armnetwork.NewScopeConnectionsClient,
		},
		{
      Name: "admin_rule_collection",
      Struct: &armnetwork.AdminRuleCollection{},
      ResponseStruct: &armnetwork.AdminRuleCollectionsClientListResponse{},
      Client: &armnetwork.AdminRuleCollectionsClient{},
      ListFunc: (&armnetwork.AdminRuleCollectionsClient{}).NewListPager,
			NewFunc: armnetwork.NewAdminRuleCollectionsClient,
		},
		{
      Name: "application_gateway",
      Struct: &armnetwork.ApplicationGateway{},
      ResponseStruct: &armnetwork.ApplicationGatewaysClientListResponse{},
      Client: &armnetwork.ApplicationGatewaysClient{},
      ListFunc: (&armnetwork.ApplicationGatewaysClient{}).NewListPager,
			NewFunc: armnetwork.NewApplicationGatewaysClient,
		},
		{
      Name: "application_security_group",
      Struct: &armnetwork.ApplicationSecurityGroup{},
      ResponseStruct: &armnetwork.ApplicationSecurityGroupsClientListResponse{},
      Client: &armnetwork.ApplicationSecurityGroupsClient{},
      ListFunc: (&armnetwork.ApplicationSecurityGroupsClient{}).NewListPager,
			NewFunc: armnetwork.NewApplicationSecurityGroupsClient,
		},
		{
      Name: "express_route_circuit_authorization",
      Struct: &armnetwork.ExpressRouteCircuitAuthorization{},
      ResponseStruct: &armnetwork.ExpressRouteCircuitAuthorizationsClientListResponse{},
      Client: &armnetwork.ExpressRouteCircuitAuthorizationsClient{},
      ListFunc: (&armnetwork.ExpressRouteCircuitAuthorizationsClient{}).NewListPager,
			NewFunc: armnetwork.NewExpressRouteCircuitAuthorizationsClient,
		},
		{
      Name: "firewall_policy_rule_collection_group",
      Struct: &armnetwork.FirewallPolicyRuleCollectionGroup{},
      ResponseStruct: &armnetwork.FirewallPolicyRuleCollectionGroupsClientListResponse{},
      Client: &armnetwork.FirewallPolicyRuleCollectionGroupsClient{},
      ListFunc: (&armnetwork.FirewallPolicyRuleCollectionGroupsClient{}).NewListPager,
			NewFunc: armnetwork.NewFirewallPolicyRuleCollectionGroupsClient,
		},
		{
      Name: "security_group",
      Struct: &armnetwork.SecurityGroup{},
      ResponseStruct: &armnetwork.SecurityGroupsClientListResponse{},
      Client: &armnetwork.SecurityGroupsClient{},
      ListFunc: (&armnetwork.SecurityGroupsClient{}).NewListPager,
			NewFunc: armnetwork.NewSecurityGroupsClient,
		},
		{
      Name: "public_ip_prefix",
      Struct: &armnetwork.PublicIPPrefix{},
      ResponseStruct: &armnetwork.PublicIPPrefixesClientListResponse{},
      Client: &armnetwork.PublicIPPrefixesClient{},
      ListFunc: (&armnetwork.PublicIPPrefixesClient{}).NewListPager,
			NewFunc: armnetwork.NewPublicIPPrefixesClient,
		},
		{
      Name: "route_filter",
      Struct: &armnetwork.RouteFilter{},
      ResponseStruct: &armnetwork.RouteFiltersClientListResponse{},
      Client: &armnetwork.RouteFiltersClient{},
      ListFunc: (&armnetwork.RouteFiltersClient{}).NewListPager,
			NewFunc: armnetwork.NewRouteFiltersClient,
		},
		{
      Name: "security_rule",
      Struct: &armnetwork.SecurityRule{},
      ResponseStruct: &armnetwork.SecurityRulesClientListResponse{},
      Client: &armnetwork.SecurityRulesClient{},
      ListFunc: (&armnetwork.SecurityRulesClient{}).NewListPager,
			NewFunc: armnetwork.NewSecurityRulesClient,
		},
		{
      Name: "vpn_gateway",
      Struct: &armnetwork.VPNGateway{},
      ResponseStruct: &armnetwork.VPNGatewaysClientListResponse{},
      Client: &armnetwork.VPNGatewaysClient{},
      ListFunc: (&armnetwork.VPNGatewaysClient{}).NewListPager,
			NewFunc: armnetwork.NewVPNGatewaysClient,
		},
		{
      Name: "web_application_firewall_policy",
      Struct: &armnetwork.WebApplicationFirewallPolicy{},
      ResponseStruct: &armnetwork.WebApplicationFirewallPoliciesClientListResponse{},
      Client: &armnetwork.WebApplicationFirewallPoliciesClient{},
      ListFunc: (&armnetwork.WebApplicationFirewallPoliciesClient{}).NewListPager,
			NewFunc: armnetwork.NewWebApplicationFirewallPoliciesClient,
		},
		{
      Name: "express_route_service_provider",
      Struct: &armnetwork.ExpressRouteServiceProvider{},
      ResponseStruct: &armnetwork.ExpressRouteServiceProvidersClientListResponse{},
      Client: &armnetwork.ExpressRouteServiceProvidersClient{},
      ListFunc: (&armnetwork.ExpressRouteServiceProvidersClient{}).NewListPager,
			NewFunc: armnetwork.NewExpressRouteServiceProvidersClient,
		},
		{
      Name: "ip_allocation",
      Struct: &armnetwork.IPAllocation{},
      ResponseStruct: &armnetwork.IPAllocationsClientListResponse{},
      Client: &armnetwork.IPAllocationsClient{},
      ListFunc: (&armnetwork.IPAllocationsClient{}).NewListPager,
			NewFunc: armnetwork.NewIPAllocationsClient,
		},
		{
      Name: "virtual_network",
      Struct: &armnetwork.VirtualNetwork{},
      ResponseStruct: &armnetwork.VirtualNetworksClientListResponse{},
      Client: &armnetwork.VirtualNetworksClient{},
      ListFunc: (&armnetwork.VirtualNetworksClient{}).NewListPager,
			NewFunc: armnetwork.NewVirtualNetworksClient,
		},
		{
      Name: "express_route_cross_connection",
      Struct: &armnetwork.ExpressRouteCrossConnection{},
      ResponseStruct: &armnetwork.ExpressRouteCrossConnectionsClientListResponse{},
      Client: &armnetwork.ExpressRouteCrossConnectionsClient{},
      ListFunc: (&armnetwork.ExpressRouteCrossConnectionsClient{}).NewListPager,
			NewFunc: armnetwork.NewExpressRouteCrossConnectionsClient,
		},
		{
      Name: "express_route_port",
      Struct: &armnetwork.ExpressRoutePort{},
      ResponseStruct: &armnetwork.ExpressRoutePortsClientListResponse{},
      Client: &armnetwork.ExpressRoutePortsClient{},
      ListFunc: (&armnetwork.ExpressRoutePortsClient{}).NewListPager,
			NewFunc: armnetwork.NewExpressRoutePortsClient,
		},
		{
      Name: "private_link_service",
      Struct: &armnetwork.PrivateLinkService{},
      ResponseStruct: &armnetwork.PrivateLinkServicesClientListResponse{},
      Client: &armnetwork.PrivateLinkServicesClient{},
      ListFunc: (&armnetwork.PrivateLinkServicesClient{}).NewListPager,
			NewFunc: armnetwork.NewPrivateLinkServicesClient,
		},
		// {
    //   Name: "base_admin_rule_classification",
    //   Struct: &armnetwork.BaseAdminRuleClassification{},
    //   ResponseStruct: &armnetwork.AdminRulesClientListResponse{},
    //   Client: &armnetwork.AdminRulesClient{},
    //   ListFunc: (&armnetwork.AdminRulesClient{}).NewListPager,
		// 	NewFunc: armnetwork.NewAdminRulesClient,
		// },
		{
      Name: "application_gateway_private_endpoint_connection",
      Struct: &armnetwork.ApplicationGatewayPrivateEndpointConnection{},
      ResponseStruct: &armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientListResponse{},
      Client: &armnetwork.ApplicationGatewayPrivateEndpointConnectionsClient{},
      ListFunc: (&armnetwork.ApplicationGatewayPrivateEndpointConnectionsClient{}).NewListPager,
			NewFunc: armnetwork.NewApplicationGatewayPrivateEndpointConnectionsClient,
		},
		{
      Name: "connectivity_configuration",
      Struct: &armnetwork.ConnectivityConfiguration{},
      ResponseStruct: &armnetwork.ConnectivityConfigurationsClientListResponse{},
      Client: &armnetwork.ConnectivityConfigurationsClient{},
      ListFunc: (&armnetwork.ConnectivityConfigurationsClient{}).NewListPager,
			NewFunc: armnetwork.NewConnectivityConfigurationsClient,
		},
		{
      Name: "express_route_port_authorization",
      Struct: &armnetwork.ExpressRoutePortAuthorization{},
      ResponseStruct: &armnetwork.ExpressRoutePortAuthorizationsClientListResponse{},
      Client: &armnetwork.ExpressRoutePortAuthorizationsClient{},
      ListFunc: (&armnetwork.ExpressRoutePortAuthorizationsClient{}).NewListPager,
			NewFunc: armnetwork.NewExpressRoutePortAuthorizationsClient,
		},
		{
      Name: "interface",
      Struct: &armnetwork.Interface{},
      ResponseStruct: &armnetwork.LoadBalancerNetworkInterfacesClientListResponse{},
      Client: &armnetwork.LoadBalancerNetworkInterfacesClient{},
      ListFunc: (&armnetwork.LoadBalancerNetworkInterfacesClient{}).NewListPager,
			NewFunc: armnetwork.NewLoadBalancerNetworkInterfacesClient,
		},
		{
      Name: "virtual_router",
      Struct: &armnetwork.VirtualRouter{},
      ResponseStruct: &armnetwork.VirtualRoutersClientListResponse{},
      Client: &armnetwork.VirtualRoutersClient{},
      ListFunc: (&armnetwork.VirtualRoutersClient{}).NewListPager,
			NewFunc: armnetwork.NewVirtualRoutersClient,
		},
		{
      Name: "group",
      Struct: &armnetwork.Group{},
      ResponseStruct: &armnetwork.GroupsClientListResponse{},
      Client: &armnetwork.GroupsClient{},
      ListFunc: (&armnetwork.GroupsClient{}).NewListPager,
			NewFunc: armnetwork.NewGroupsClient,
		},
		{
      Name: "interface",
      Struct: &armnetwork.Interface{},
      ResponseStruct: &armnetwork.InterfacesClientListResponse{},
      Client: &armnetwork.InterfacesClient{},
      ListFunc: (&armnetwork.InterfacesClient{}).NewListPager,
			NewFunc: armnetwork.NewInterfacesClient,
		},
		{
      Name: "probe",
      Struct: &armnetwork.Probe{},
      ResponseStruct: &armnetwork.LoadBalancerProbesClientListResponse{},
      Client: &armnetwork.LoadBalancerProbesClient{},
      ListFunc: (&armnetwork.LoadBalancerProbesClient{}).NewListPager,
			NewFunc: armnetwork.NewLoadBalancerProbesClient,
		},
		{
      Name: "profile",
      Struct: &armnetwork.Profile{},
      ResponseStruct: &armnetwork.ProfilesClientListResponse{},
      Client: &armnetwork.ProfilesClient{},
      ListFunc: (&armnetwork.ProfilesClient{}).NewListPager,
			NewFunc: armnetwork.NewProfilesClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armnetwork"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armnetwork()...)
}