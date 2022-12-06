# Table: aws_kinesis_streams

https://docs.aws.amazon.com/kinesis/latest/APIReference/API_StreamDescriptionSummary.html

The primary key for this table is **arn**.



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