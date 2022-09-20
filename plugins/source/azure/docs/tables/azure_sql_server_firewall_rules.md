
# Table: azure_sql_server_firewall_rules
The list of server firewall rules.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_cq_id|uuid|Unique ID of azure_sql_servers table (FK)|
|kind|text|Kind of server that contains this firewall rule|
|location|text|Location of the server that contains this firewall rule|
|start_ip_address|text|The start IP address of the firewall rule.|
|end_ip_address|text|The end IP address of the firewall rule. Must be IPv4.|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
