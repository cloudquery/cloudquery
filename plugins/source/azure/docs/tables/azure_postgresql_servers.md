# Table: azure_postgresql_servers

https://learn.microsoft.com/en-us/rest/api/postgresql/singleserver/servers/list?tabs=HTTP#server

The primary key for this table is **id**.

## Relations

The following tables depend on azure_postgresql_servers:
  - [azure_postgresql_databases](azure_postgresql_databases.md)
  - [azure_postgresql_server_configurations](azure_postgresql_server_configurations.md)
  - [azure_postgresql_server_firewall_rules](azure_postgresql_server_firewall_rules.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|identity|JSON|
|properties|JSON|
|sku|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|