
# Table: azure_iothub_hub_private_endpoint_connections
PrivateEndpointConnection the private endpoint connection of an IotHub
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hub_cq_id|uuid|Unique CloudQuery ID of azure_iothub_hubs table (FK)|
|id|text|The resource identifier.|
|name|text|The resource name.|
|type|text|The resource type.|
|private_endpoint_id|text|The resource identifier.|
|status|text|The status of a private endpoint connection|
|description|text|The description for the current state of a private endpoint connection|
|actions_required|text|Actions required for a private endpoint connection|
