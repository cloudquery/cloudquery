# Table: aws_dynamodbstreams_streams

This table shows data for Amazon DynamoDB Streams.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_StreamDescription.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|creation_request_date_time|`timestamp[us, tz=UTC]`|
|key_schema|`json`|
|last_evaluated_shard_id|`utf8`|
|shards|`json`|
|stream_arn|`utf8`|
|stream_label|`utf8`|
|stream_status|`utf8`|
|stream_view_type|`utf8`|
|table_name|`utf8`|