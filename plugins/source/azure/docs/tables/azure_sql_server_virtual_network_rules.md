
# Table: azure_sql_server_virtual_network_rules
List of virtual network for a server
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_cq_id|uuid|Unique ID of azure_sql_servers table (FK)|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|The virtual network rule type|
|subnet_id|text|The ARM resource id of the virtual network subnet.|
|ignore_missing_vnet_service_endpoint|boolean|Create firewall rule before the virtual network has vnet service endpoint enabled.|
|state|text|Virtual Network Rule State|
