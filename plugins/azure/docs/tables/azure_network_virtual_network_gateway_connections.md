
# Table: azure_network_virtual_network_gateway_connections
Azure virtual network gateway connection
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_network_gateway_cq_id|uuid|Unique CloudQuery ID of azure_network_virtual_network_gateways table (FK)|
|authorization_key|text|The authorizationKey.|
|virtual_network_gateway_1|jsonb|The reference to virtual network gateway resource.|
|virtual_network_gateway_2|jsonb|The reference to virtual network gateway resource.|
|local_network_gateway_2|jsonb|The reference to local network gateway resource.|
|connection_type|text|Gateway connection type.|
|connection_protocol|text|Connection protocol used for this connection.|
|routing_weight|integer|The routing weight.|
|connection_mode|text|The connection mode for this connection.|
|shared_key|text|The IPSec shared key.|
|connection_status|text|Virtual Network Gateway connection status.|
|tunnel_connection_status|jsonb|Collection of all tunnels' connection health status.|
|egress_bytes_transferred|bigint|The egress bytes transferred in this connection.|
|ingress_bytes_transferred|bigint|The ingress bytes transferred in this connection.|
|peer_id|text|The reference to peerings resource.|
|enable_bgp|boolean|EnableBgp flag.|
|use_policy_based_traffic_selectors|boolean|Enable policy-based traffic selectors.|
|ipsec_policies|jsonb|The IPSec Policies to be considered by this connection.|
|traffic_selector_policies|jsonb|The Traffic Selector Policies to be considered by this connection.|
|resource_guid|text|The resource GUID property of the virtual network gateway connection resource.|
|provisioning_state|text|The provisioning state of the virtual network gateway connection resource.|
|express_route_gateway_bypass|boolean|Bypass ExpressRoute Gateway for data forwarding.|
|etag|text|A unique read-only string that changes whenever the resource is updated.|
|id|text|Resource ID.|
|name|text|Resource name.|
|type|text|Resource type.|
|location|text|Resource location.|
|tags|jsonb|Resource tags.|
