
# Table: azure_postgresql_server_firewall_rules
Azure postgresql server firewall rule
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_cq_id|uuid|Unique ID of azure_postgresql_servers table (FK)|
|start_ip_address|text|The start IP address of the server firewall rule Must be IPv4 format|
|end_ip_address|text|The end IP address of the server firewall rule Must be IPv4 format|
|id|text|Fully qualified resource ID for the resource Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}|
|name|text|The name of the resource|
|type|text|The type of the resource Eg "MicrosoftCompute/virtualMachines" or "MicrosoftStorage/storageAccounts"|
