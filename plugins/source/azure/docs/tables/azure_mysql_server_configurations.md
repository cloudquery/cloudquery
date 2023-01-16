# Table: azure_mysql_server_configurations

The primary key for this table is **id**.

## Relations

This table depends on [azure_mysql_servers](azure_mysql_servers.md).

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