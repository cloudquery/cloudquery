# Table: azure_mysql_servers

The primary key for this table is **id**.

## Relations

The following tables depend on azure_mysql_servers:
  - [azure_mysql_server_configurations](azure_mysql_server_configurations.md)

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