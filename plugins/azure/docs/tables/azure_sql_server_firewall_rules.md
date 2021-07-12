
# Table: azure_sql_server_firewall_rules
The list of server firewall rules.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_cq_id|uuid|Unique ID of azure_sql_servers table (FK)|
|kind|text|Kind of server that contains this firewall rule|
|location|text|Location of the server that contains this firewall rule|
|start_ip_address|text|The start IP address of the firewall rule Must be IPv4 format Use value '0000' to represent all Azure-internal IP addresses|
|end_ip_address|text|The end IP address of the firewall rule Must be IPv4 format Must be greater than or equal to startIpAddress Use value '0000' to represent all Azure-internal IP addresses|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
