# Table: azure_sql_server_virtual_network_rules

This table shows data for Azure SQL Server Virtual Network Rules.

https://learn.microsoft.com/en-us/rest/api/sql/2020-08-01-preview/virtual-network-rules/list-by-server?tabs=HTTP#virtualnetworkrule

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_servers](azure_sql_servers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|