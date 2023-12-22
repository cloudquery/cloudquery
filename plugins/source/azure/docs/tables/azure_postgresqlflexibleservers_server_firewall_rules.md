# Table: azure_postgresqlflexibleservers_server_firewall_rules

This table shows data for Azure PostgreSQL Flexible Servers Server Firewall Rules.

https://learn.microsoft.com/en-us/rest/api/postgresql/flexibleserver/firewall-rules/list-by-server

The primary key for this table is **id**.

## Relations

This table depends on [azure_postgresqlflexibleservers_servers](azure_postgresqlflexibleservers_servers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|