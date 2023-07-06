# Table: aws_glue_job_runs

This table shows data for Glue Job Runs.

https://docs.aws.amazon.com/glue/latest/webapi/API_JobRun.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_glue_jobs](aws_glue_jobs).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|job_arn|`utf8`|
|allocated_capacity|`int64`|
|arguments|`json`|
|attempt|`int64`|
|completed_on|`timestamp[us, tz=UTC]`|
|dpu_seconds|`float64`|
|error_message|`utf8`|
|execution_class|`utf8`|
|execution_time|`int64`|
|glue_version|`utf8`|
|id|`utf8`|
|job_name|`utf8`|
|job_run_state|`utf8`|
|last_modified_on|`timestamp[us, tz=UTC]`|
|log_group_name|`utf8`|
|max_capacity|`float64`|
|notification_property|`json`|
|number_of_workers|`int64`|
|predecessor_runs|`json`|
|previous_run_id|`utf8`|
|security_configuration|`utf8`|
|started_on|`timestamp[us, tz=UTC]`|
|timeout|`int64`|
|trigger_name|`utf8`|
|worker_type|`utf8`|