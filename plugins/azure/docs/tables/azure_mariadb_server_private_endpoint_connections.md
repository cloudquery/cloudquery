
# Table: azure_mariadb_server_private_endpoint_connections
List of private endpoint connections on a server
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_cq_id|uuid|Unique ID of azure_mysql_servers table (FK)|
|id|text|Resource Id of the private endpoint connection.|
|status|text|The private link service connection status.|
|status_description|text|The private link service connection description.|
|provisioning_state|text|State of the private endpoint connection.|
