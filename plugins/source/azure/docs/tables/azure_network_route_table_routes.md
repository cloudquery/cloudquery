
# Table: azure_network_route_table_routes
Azure route table route
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|route_table_cq_id|uuid|Unique CloudQuery ID of azure_network_route_tables table (FK)|
|id|text|Resource ID.|
|address_prefix|text|The destination CIDR to which the route applies.|
|next_hop_type|text|The type of Azure hop the packet should be sent to.|
|next_hop_ip_address|text|The IP address packets should be forwarded to.|
|provisioning_state|text|The provisioning state of the route resource.|
|has_bgp_override|boolean|A value indicating whether this route overrides overlapping BGP routes regardless of LPM.|
|etag|text|A unique read-only string that changes whenever the resource is updated.|
|name|text|Resource name.|
|type|text|Resource type.|
