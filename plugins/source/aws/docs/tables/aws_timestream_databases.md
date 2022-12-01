# Table: aws_timestream_databases



The primary key for this table is **arn**.

## Relations

The following tables depend on aws_timestream_databases:
  - [aws_timestream_tables](aws_timestream_tables.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|arn (PK)|String|
|creation_time|Timestamp|
|database_name|String|
|kms_key_id|String|
|last_updated_time|Timestamp|
|table_count|Int|