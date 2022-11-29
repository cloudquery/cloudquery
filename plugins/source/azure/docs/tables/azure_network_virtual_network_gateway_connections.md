# Table: azure_network_virtual_network_gateway_connections

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network#VirtualNetworkGatewayConnectionListEntity

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
|network_virtual_network_gateway_id|String|
|authorization_key|String|
|virtual_network_gateway1|JSON|
|virtual_network_gateway2|JSON|
|local_network_gateway2|JSON|
|connection_type|String|
|connection_protocol|String|
|routing_weight|Int|
|connection_mode|String|
|shared_key|String|
|connection_status|String|
|tunnel_connection_status|JSON|
|egress_bytes_transferred|Int|
|ingress_bytes_transferred|Int|
|peer|JSON|
|enable_bgp|Bool|
|use_policy_based_traffic_selectors|Bool|
|ipsec_policies|JSON|
|traffic_selector_policies|JSON|
|resource_guid|String|
|provisioning_state|String|
|express_route_gateway_bypass|Bool|
|etag|String|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|