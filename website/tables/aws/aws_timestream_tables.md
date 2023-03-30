# Table: aws_timestream_tables

This table shows data for Timestream Tables.

https://docs.aws.amazon.com/timestream/latest/developerguide/API_Table.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_timestream_databases](aws_timestream_databases).

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