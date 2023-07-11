# Table: aws_glue_ml_transform_task_runs

This table shows data for Glue ML Transform Task Runs.

https://docs.aws.amazon.com/glue/latest/webapi/API_TaskRun.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_glue_ml_transforms](aws_glue_ml_transforms).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|ml_transform_arn|`utf8`|
|completed_on|`timestamp[us, tz=UTC]`|
|error_string|`utf8`|
|execution_time|`int64`|
|last_modified_on|`timestamp[us, tz=UTC]`|
|log_group_name|`utf8`|
|properties|`json`|
|started_on|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|task_run_id|`utf8`|
|transform_id|`utf8`|