# Table: aws_glue_ml_transform_task_runs



The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_glue_ml_transforms](aws_glue_ml_transforms.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|ml_transform_arn|String|
|completed_on|Timestamp|
|error_string|String|
|execution_time|Int|
|last_modified_on|Timestamp|
|log_group_name|String|
|properties|JSON|
|started_on|Timestamp|
|status|String|
|task_run_id|String|
|transform_id|String|