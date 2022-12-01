# Table: aws_glue_ml_transforms



The primary key for this table is **arn**.

## Relations

The following tables depend on aws_glue_ml_transforms:
  - [aws_glue_ml_transform_task_runs](aws_glue_ml_transform_task_runs.md)

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
|schema|JSON|
|created_on|Timestamp|
|description|String|
|evaluation_metrics|JSON|
|glue_version|String|
|input_record_tables|JSON|
|label_count|Int|
|last_modified_on|Timestamp|
|max_capacity|Float|
|max_retries|Int|
|name|String|
|number_of_workers|Int|
|parameters|JSON|
|role|String|
|status|String|
|timeout|Int|
|transform_encryption|JSON|
|transform_id|String|
|worker_type|String|