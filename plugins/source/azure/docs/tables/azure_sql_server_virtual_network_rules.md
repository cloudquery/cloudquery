# Table: azure_sql_server_virtual_network_rules

https://learn.microsoft.com/en-us/rest/api/sql/2020-08-01-preview/virtual-network-rules/list-by-server?tabs=HTTP#virtualnetworkrule

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
|properties|JSON|
|id (PK)|String|
|name|String|
|type|String|