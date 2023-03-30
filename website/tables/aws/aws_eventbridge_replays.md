# Table: aws_eventbridge_replays

This table shows data for Amazon EventBridge Replays.

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_DescribeReplay.html

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
|description|String|
|destination|JSON|
|event_end_time|Timestamp|
|event_last_replayed_time|Timestamp|
|event_source_arn|String|
|event_start_time|Timestamp|
|replay_arn|String|
|replay_end_time|Timestamp|
|replay_name|String|
|replay_start_time|Timestamp|
|state|String|
|state_reason|String|