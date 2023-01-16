# Table: azure_sql_servers

The primary key for this table is **id**.

## Relations

The following tables depend on azure_sql_servers:
  - [azure_sql_virtual_network_rules](azure_sql_virtual_network_rules.md)

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
|tags|JSON|
|kind|String|
|name|String|
|type|String|