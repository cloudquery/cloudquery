
# Table: azure_servicebus_namespace_private_endpoint_connections
List of private endpoint connections.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|namespace_cq_id|uuid|Unique ID of azure_servicebus_namespaces table (FK)|
|id|text|Resource Id.|
|name|text|Resource name.|
|type|text|Resource type.|
|system_data|jsonb|The system meta data relating to this resource.|
|status|text|The private link service connection status.|
|status_description|text|The private link service connection description.|
|provisioning_state|text|State of the private endpoint connection.|
