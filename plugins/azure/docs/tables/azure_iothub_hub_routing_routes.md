
# Table: azure_iothub_hub_routing_routes
RouteProperties the properties of a routing rule that your IoT hub uses to route messages to endpoints.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hub_cq_id|uuid|Unique CloudQuery ID of azure_iothub_hubs table (FK)|
|name|text|The name of the route|
|source|text|The source that the routing rule is to be applied to, such as DeviceMessages|
|condition|text|The condition that is evaluated to apply the routing rule|
|endpoint_names|text[]|The list of endpoints to which messages that satisfy the condition are routed|
|is_enabled|boolean|Used to specify whether a route is enabled.|
