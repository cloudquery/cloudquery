# Table: aws_scheduler_schedules

This table shows data for Amazon EventBridge Scheduler Schedules.

https://docs.aws.amazon.com/scheduler/latest/APIReference/API_GetScheduleOutput.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|creation_date|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|end_date|`timestamp[us, tz=UTC]`|
|flexible_time_window|`json`|
|group_name|`utf8`|
|kms_key_arn|`utf8`|
|last_modification_date|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|schedule_expression|`utf8`|
|schedule_expression_timezone|`utf8`|
|start_date|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|target|`json`|