# Table: azure_postgresql_server_firewall_rules

https://learn.microsoft.com/en-us/rest/api/postgresql/singleserver/firewall-rules/list-by-server?tabs=HTTP#firewallrule

The primary key for this table is **id**.

## Relations

This table depends on [azure_postgresql_servers](azure_postgresql_servers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|properties|JSON|
|name|String|
|type|String|