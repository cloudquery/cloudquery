# Table: aws_glue_job_runs



The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_glue_jobs](aws_glue_jobs.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|job_arn|String|
|allocated_capacity|Int|
|arguments|JSON|
|attempt|Int|
|completed_on|Timestamp|
|dpu_seconds|Float|
|error_message|String|
|execution_class|String|
|execution_time|Int|
|glue_version|String|
|id|String|
|job_name|String|
|job_run_state|String|
|last_modified_on|Timestamp|
|log_group_name|String|
|max_capacity|Float|
|notification_property|JSON|
|number_of_workers|Int|
|predecessor_runs|JSON|
|previous_run_id|String|
|security_configuration|String|
|started_on|Timestamp|
|timeout|Int|
|trigger_name|String|
|worker_type|String|