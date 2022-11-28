# Table: aws_glue_jobs



The primary key for this table is **arn**.

## Relations

The following tables depend on aws_glue_jobs:
  - [aws_glue_job_runs](aws_glue_job_runs.md)

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
|allocated_capacity|Int|
|code_gen_configuration_nodes|JSON|
|command|JSON|
|connections|JSON|
|created_on|Timestamp|
|default_arguments|JSON|
|description|String|
|execution_class|String|
|execution_property|JSON|
|glue_version|String|
|last_modified_on|Timestamp|
|log_uri|String|
|max_capacity|Float|
|max_retries|Int|
|name|String|
|non_overridable_arguments|JSON|
|notification_property|JSON|
|number_of_workers|Int|
|role|String|
|security_configuration|String|
|source_control_details|JSON|
|timeout|Int|
|worker_type|String|