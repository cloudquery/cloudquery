# Table: aws_scheduler_schedules

https://docs.aws.amazon.com/scheduler/latest/APIReference/API_GetScheduleOutput.html

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
|tags|JSON|
|arn (PK)|String|
|creation_date|Timestamp|
|description|String|
|end_date|Timestamp|
|flexible_time_window|JSON|
|group_name|String|
|kms_key_arn|String|
|last_modification_date|Timestamp|
|name|String|
|schedule_expression|String|
|schedule_expression_timezone|String|
|start_date|Timestamp|
|state|String|
|target|JSON|