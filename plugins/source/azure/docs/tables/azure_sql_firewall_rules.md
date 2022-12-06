# Table: azure_sql_firewall_rules

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql#FirewallRule

The primary key for this table is **id**.

## Relations
This table depends on [azure_sql_servers](azure_sql_servers.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|sql_server_id|String|
|kind|String|
|location|String|
|start_ip_address|String|
|end_ip_address|String|
|id (PK)|String|
|name|String|
|type|String|