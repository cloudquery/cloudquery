# Table: azure_network_virtual_network_gateways

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network#VirtualNetworkGateway

The primary key for this table is **id**.

## Relations
This table depends on [azure_network_virtual_networks](azure_network_virtual_networks.md).

The following tables depend on azure_network_virtual_network_gateways:
  - [azure_network_virtual_network_gateway_connections](azure_network_virtual_network_gateway_connections.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|network_virtual_network_id|String|
|extended_location|JSON|
|ip_configurations|JSON|
|gateway_type|String|
|vpn_type|String|
|vpn_gateway_generation|String|
|enable_bgp|Bool|
|enable_private_ip_address|Bool|
|active_active|Bool|
|gateway_default_site|JSON|
|sku|JSON|
|vpn_client_configuration|JSON|
|bgp_settings|JSON|
|custom_routes|JSON|
|resource_guid|String|
|provisioning_state|String|
|enable_dns_forwarding|Bool|
|inbound_dns_forwarding_endpoint|String|
|v_net_extended_location_resource_id|String|
|etag|String|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|