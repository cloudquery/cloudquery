# Table: azure_postgresql_server_firewall_rules

This table shows data for Azure PostgreSQL Server Firewall Rules.

https://learn.microsoft.com/en-us/rest/api/postgresql/singleserver/firewall-rules/list-by-server?tabs=HTTP#firewallrule

The primary key for this table is **id**.

## Relations

This table depends on [azure_postgresql_servers](azure_postgresql_servers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|