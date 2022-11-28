# Table: aws_timestream_tables



The primary key for this table is **arn**.

## Relations
This table depends on [aws_timestream_databases](aws_timestream_databases.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|creation_time|Timestamp|
|database_name|String|
|last_updated_time|Timestamp|
|magnetic_store_write_properties|JSON|
|retention_properties|JSON|
|table_name|String|
|table_status|String|