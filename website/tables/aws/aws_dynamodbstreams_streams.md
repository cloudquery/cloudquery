# Table: aws_dynamodbstreams_streams

This table shows data for Amazon DynamoDB Streams.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_StreamDescription.html

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
|creation_request_date_time|Timestamp|
|key_schema|JSON|
|last_evaluated_shard_id|String|
|shards|JSON|
|stream_arn|String|
|stream_label|String|
|stream_status|String|
|stream_view_type|String|
|table_name|String|