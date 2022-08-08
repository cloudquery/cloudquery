
# Table: azure_network_virtual_network_gateways
Azure virtual network gateway
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_network_cq_id|uuid|Unique CloudQuery ID of azure_network_virtual_networks table (FK)|
|extended_location_name|text|The name of the extended location.|
|extended_location_type|text|The type of the extended location.|
|ip_configurations|jsonb|IP configurations for virtual network gateway.|
|gateway_type|text|The type of this virtual network gateway.|
|vpn_type|text|The type of this virtual network gateway.|
|vpn_gateway_generation|text|The generation for this VirtualNetworkGateway.|
|enable_bgp|boolean|Whether BGP is enabled for this virtual network gateway or not.|
|enable_private_ip_address|boolean|Whether private IP needs to be enabled on this gateway for connections or not.|
|active_active|boolean|ActiveActive flag.|
|gateway_default_site_id|text|The reference to the LocalNetworkGateway resource which represents local network site having default routes.|
|sku_name|text|Gateway SKU name.|
|sku_tier|text|Gateway SKU tier.|
|sku_capacity|integer|READ-ONLY; The capacity.|
|vpn_client_configuration_address_pool|text[]|The reference to the address space resource which represents Address space for P2S VpnClient.|
|vpn_client_configuration_root_certificates|jsonb|VpnClientRootCertificate for virtual network gateway.|
|vpn_client_configuration_revoked_certificates|jsonb|VpnClientRevokedCertificate for Virtual network gateway.|
|vpn_client_configuration_protocols|text[]|VpnClientProtocols for Virtual network gateway.|
|vpn_client_configuration_authentication_types|text[]|VPN authentication types for the virtual network gateway.|
|vpn_client_configuration_ipsec_policies|jsonb|VpnClientIpsecPolicies for virtual network gateway P2S client.|
|vpn_client_configuration_radius_server_address|text|The radius server address property of the VirtualNetworkGateway resource for vpn client connection.|
|vpn_client_configuration_radius_server_secret|text|The radius secret property of the VirtualNetworkGateway resource for vpn client connection.|
|vpn_client_configuration_radius_servers|jsonb|The radiusServers property for multiple radius server configuration.|
|vpn_client_configuration_aad_tenant|text|The AADTenant property of the VirtualNetworkGateway resource for vpn client connection used for AAD authentication.|
|vpn_client_configuration_aad_audience|text|The AADAudience property of the VirtualNetworkGateway resource for vpn client connection used for AAD authentication.|
|vpn_client_configuration_aad_issuer|text|The AADIssuer property of the VirtualNetworkGateway resource for vpn client connection used for AAD authentication.|
|bgp_settings_asn|bigint|The BGP speaker's ASN.|
|bgp_settings_bgp_peering_address|text|The BGP peering address and BGP identifier of this BGP speaker.|
|bgp_settings_peer_weight|integer|The weight added to routes learned from this BGP speaker.|
|bgp_settings_bgp_peering_addresses|jsonb|BGP peering address with IP configuration ID for virtual network gateway.|
|custom_routes_address_prefixes|text[]|The reference to the address space resource which represents the custom routes address space specified by the customer for virtual network gateway and VpnClient.|
|resource_guid|text|The resource GUID property of the virtual network gateway resource.|
|provisioning_state|text|The provisioning state of the virtual network gateway resource.|
|enable_dns_forwarding|boolean|Whether dns forwarding is enabled or not.|
|inbound_dns_forwarding_endpoint|text|The IP address allocated by the gateway to which dns requests can be sent.|
|vnet_extended_location_resource_id|text|Customer vnet resource id.|
|etag|text|A unique read-only string that changes whenever the resource is updated.|
|id|text|Resource ID.|
|name|text|Resource name.|
|type|text|Resource type.|
|location|text|Resource location.|
|tags|jsonb|Resource tags.|
