
# Table: azure_network_express_route_connections
ExpressRouteConnection resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|express_route_gateway_cq_id|uuid|Unique CloudQuery ID of azure_network_express_route_gateways table (FK)|
|id|text|Resource ID.|
|authorization_key|text|Authorization key to establish the connection.|
|enable_internet_security|boolean|Enable internet security.|
|express_route_circuit_peering_id|text|The ID of the ExpressRoute circuit peering.|
|name|text|Resource name.|
|provisioning_state|text|The provisioning state of the express route connection resource.|
|routing_weight|integer|The routing weight associated to the connection.|
