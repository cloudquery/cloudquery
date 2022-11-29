# Table: aws_lightsail_database_log_events



The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_lightsail_databases](aws_lightsail_databases.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|database_arn|String|
|created_at|Timestamp|
|message|String|
|log_stream_name|String|