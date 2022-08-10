
# Table: azure_network_express_route_circuit_connections
Express Route Circuit Connection in an ExpressRouteCircuitPeering resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|express_route_circuit_peering_cq_id|uuid|Unique CloudQuery ID of azure_network_express_route_circuit_peerings table (FK)|
|id|text|Resource ID.|
|address_prefix|text|/29 IP address space to carve out Customer addresses for tunnels.|
|circuit_connection_status|text|Express Route Circuit connection state.|
|etag|text|A unique read-only string that changes whenever the resource is updated.|
|express_route_circuit_peering_id|text|Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.|
|ipv6_circuit_connection_config_address_prefix|text|/125 IP address space to carve out customer addresses for global reach.|
|ipv6_circuit_connection_config_circuit_connection_status|text|Express Route Circuit connection state.|
|name|text|Resource name.|
|peer_express_route_circuit_peering_id|text|Reference to Express Route Circuit Private Peering Resource of the peered circuit.|
|provisioning_state|text|The provisioning state of the express route circuit connection resource.|
|type|text|Resource type.|
