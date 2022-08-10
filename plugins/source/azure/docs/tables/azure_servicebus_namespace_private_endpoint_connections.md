
# Table: azure_servicebus_namespace_private_endpoint_connections
List of private endpoint connections
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|namespace_cq_id|uuid|Unique CloudQuery ID of azure_servicebus_namespaces table (FK)|
|private_endpoint_id|text|The ARM identifier for Private Endpoint|
|status|text|Status of the connection|
|status_description|text|Description of the connection state|
|provisioning_state|text|Provisioning state of the Private Endpoint Connection|
|system_data|jsonb|The system meta data relating to this resource|
|id|text|Resource Id|
|name|text|Resource name|
|type|text|Resource type|
