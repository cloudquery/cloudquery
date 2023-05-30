# Table: azure_postgresql_servers

This table shows data for Azure PostgreSQL Servers.

https://learn.microsoft.com/en-us/rest/api/postgresql/singleserver/servers/list?tabs=HTTP#server

The primary key for this table is **id**.

## Relations

The following tables depend on azure_postgresql_servers:
  - [azure_postgresql_databases](azure_postgresql_databases)
  - [azure_postgresql_server_configurations](azure_postgresql_server_configurations)
  - [azure_postgresql_server_firewall_rules](azure_postgresql_server_firewall_rules)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|identity|`json`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|