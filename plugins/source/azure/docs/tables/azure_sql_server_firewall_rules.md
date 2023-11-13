# Table: azure_sql_server_firewall_rules

This table shows data for Azure SQL Server Firewall Rules.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/firewall-rules/list-by-server?tabs=HTTP#firewallrule

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_servers](azure_sql_servers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|name|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|type|`utf8`|