# Table: azure_network_virtual_network_gateway_connections

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2#VirtualNetworkGatewayConnection

The primary key for this table is **id**.

## Relations
This table depends on [azure_network_virtual_network_gateways](azure_network_virtual_network_gateways.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|connection_type|String|
|virtual_network_gateway1|JSON|
|authorization_key|String|
|connection_mode|String|
|connection_protocol|String|
|dpd_timeout_seconds|Int|
|egress_nat_rules|JSON|
|enable_bgp|Bool|
|enable_private_link_fast_path|Bool|
|express_route_gateway_bypass|Bool|
|gateway_custom_bgp_ip_addresses|JSON|
|ipsec_policies|JSON|
|ingress_nat_rules|JSON|
|local_network_gateway2|JSON|
|peer|JSON|
|routing_weight|Int|
|shared_key|String|
|traffic_selector_policies|JSON|
|use_local_azure_ip_address|Bool|
|use_policy_based_traffic_selectors|Bool|
|virtual_network_gateway2|JSON|
|connection_status|String|
|egress_bytes_transferred|Int|
|ingress_bytes_transferred|Int|
|provisioning_state|String|
|resource_guid|String|
|tunnel_connection_status|JSON|
|id (PK)|String|
|location|String|
|tags|JSON|
|etag|String|
|name|String|
|type|String|
|virtual_network_gateway_id|String|