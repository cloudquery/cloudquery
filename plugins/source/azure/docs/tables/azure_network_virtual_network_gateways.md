# Table: azure_network_virtual_network_gateways

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2#VirtualNetworkGateway

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
|active_active|Bool|
|allow_remote_vnet_traffic|Bool|
|allow_virtual_wan_traffic|Bool|
|bgp_settings|JSON|
|custom_routes|JSON|
|disable_ip_sec_replay_protection|Bool|
|enable_bgp|Bool|
|enable_bgp_route_translation_for_nat|Bool|
|enable_dns_forwarding|Bool|
|enable_private_ip_address|Bool|
|gateway_default_site|JSON|
|gateway_type|String|
|ip_configurations|JSON|
|nat_rules|JSON|
|sku|JSON|
|v_net_extended_location_resource_id|String|
|vpn_client_configuration|JSON|
|vpn_gateway_generation|String|
|vpn_type|String|
|virtual_network_gateway_policy_groups|JSON|
|inbound_dns_forwarding_endpoint|String|
|provisioning_state|String|
|resource_guid|String|
|extended_location|JSON|
|id (PK)|String|
|location|String|
|tags|JSON|
|etag|String|
|name|String|
|type|String|
|virtual_network_id|String|