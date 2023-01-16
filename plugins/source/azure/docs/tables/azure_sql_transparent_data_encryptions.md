# Table: azure_sql_transparent_data_encryptions

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_server_databases](azure_sql_server_databases.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|properties|JSON|
|id (PK)|String|
|name|String|
|type|String|