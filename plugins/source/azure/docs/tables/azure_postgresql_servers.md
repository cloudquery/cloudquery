# Table: azure_postgresql_servers

The primary key for this table is **id**.

## Relations

The following tables depend on azure_postgresql_servers:
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
|id (PK)|String|
|location|String|
|identity|JSON|
|properties|JSON|
|sku|JSON|
|tags|JSON|
|name|String|
|type|String|