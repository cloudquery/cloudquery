
# Table: azure_cosmosdb_account_private_endpoint_connections
PrivateEndpointConnection a private endpoint connection
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_cq_id|uuid|Unique CloudQuery ID of azure_cosmosdb_accounts table (FK)|
|private_endpoint_id|text|Resource id of the private endpoint.|
|status|text|The private link service connection status.|
|actions_required|text|Any action that is required beyond basic workflow (approve/ reject/ disconnect)|
|description|text|The private link service connection description.|
|group_id|text|Group id of the private endpoint.|
|provisioning_state|text|Provisioning state of the private endpoint.|
|id|text|Fully qualified resource ID for the resource|
|name|text|The name of the resource|
|type|text|The type of the resource|
