# Table: aws_kinesis_streams


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|enhanced_monitoring|JSON|
|open_shard_count|Int|
|retention_period_hours|Int|
|stream_creation_timestamp|Timestamp|
|stream_name|String|
|stream_status|String|
|consumer_count|Int|
|encryption_type|String|
|key_id|String|
|stream_mode_details|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|