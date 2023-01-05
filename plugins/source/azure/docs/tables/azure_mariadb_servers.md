# Table: azure_mariadb_servers

The primary key for this table is **id**.

## Relations

The following tables depend on azure_mariadb_servers:
  - [azure_mariadb_server_configurations](azure_mariadb_server_configurations.md)

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
|properties|JSON|
|sku|JSON|
|tags|JSON|
|name|String|
|type|String|