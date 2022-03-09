
# Table: azure_search_service_private_endpoint_connections
PrivateEndpointConnection describes an existing Private Endpoint connection to the Azure Cognitive Search service.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of azure_search_services table (FK)|
|private_endpoint_id|text|The resource id of the private endpoint resource from Microsoft.Network provider.|
|private_link_connection_status|text|Status of the the private link service connection|
|private_link_connection_description|text|The description for the private link service connection state.|
|private_link_connection_actions_required|text|A description of any extra actions that may be required.|
|id|text|Fully qualified resource ID for the resource|
|name|text|The name of the resource|
|type|text|The type of the resource|
