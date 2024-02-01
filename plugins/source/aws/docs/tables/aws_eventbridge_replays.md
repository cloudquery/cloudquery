# Table: aws_eventbridge_replays

This table shows data for Amazon EventBridge Replays.

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_DescribeReplay.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|description|`utf8`|
|destination|`json`|
|event_end_time|`timestamp[us, tz=UTC]`|
|event_last_replayed_time|`timestamp[us, tz=UTC]`|
|event_source_arn|`utf8`|
|event_start_time|`timestamp[us, tz=UTC]`|
|replay_arn|`utf8`|
|replay_end_time|`timestamp[us, tz=UTC]`|
|replay_name|`utf8`|
|replay_start_time|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|state_reason|`utf8`|