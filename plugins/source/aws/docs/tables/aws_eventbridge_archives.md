# Table: aws_eventbridge_archives

This table shows data for Amazon EventBridge Archives.

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Archive.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|archive_name|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|event_count|`int64`|
|event_source_arn|`utf8`|
|retention_days|`int64`|
|size_bytes|`int64`|
|state|`utf8`|
|state_reason|`utf8`|