# Table: azure_sql_database_long_term_retention_policies

The primary key for this table is **_cq_id**.

## Relations

This table depends on [azure_sql_server_databases](azure_sql_server_databases.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|properties|JSON|
|id|String|
|name|String|
|type|String|