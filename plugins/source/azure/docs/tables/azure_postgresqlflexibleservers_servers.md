# Table: azure_postgresqlflexibleservers_servers

This table shows data for Azure PostgreSQL Flexible Servers Servers.

https://learn.microsoft.com/en-us/rest/api/postgresql/flexibleserver/servers/list?tabs=HTTP#server

The primary key for this table is **id**.

## Relations

The following tables depend on azure_postgresqlflexibleservers_servers:
  - [azure_postgresqlflexibleservers_server_configurations](azure_postgresqlflexibleservers_server_configurations.md)
  - [azure_postgresqlflexibleservers_server_firewall_rules](azure_postgresqlflexibleservers_server_firewall_rules.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
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
|system_data|`json`|
|type|`utf8`|