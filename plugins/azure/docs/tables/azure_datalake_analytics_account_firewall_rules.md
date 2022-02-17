
# Table: azure_datalake_analytics_account_firewall_rules
FirewallRule data Lake Analytics firewall rule information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|analytics_account_cq_id|uuid|Unique CloudQuery ID of azure_datalake_analytics_accounts table (FK)|
|start_ip_address|inet|The start IP address for the firewall rule|
|end_ip_address|inet|The end IP address for the firewall rule|
|id|text|The resource identifier|
|name|text|The resource name|
|type|text|The resource type|
