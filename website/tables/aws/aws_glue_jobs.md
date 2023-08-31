# Table: aws_glue_jobs

This table shows data for Glue Jobs.

https://docs.aws.amazon.com/glue/latest/webapi/API_Job.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_glue_jobs:
  - [aws_glue_job_runs](aws_glue_job_runs)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|allocated_capacity|`int64`|
|code_gen_configuration_nodes|`json`|
|command|`json`|
|connections|`json`|
|created_on|`timestamp[us, tz=UTC]`|
|default_arguments|`json`|
|description|`utf8`|
|execution_class|`utf8`|
|execution_property|`json`|
|glue_version|`utf8`|
|last_modified_on|`timestamp[us, tz=UTC]`|
|log_uri|`utf8`|
|max_capacity|`float64`|
|max_retries|`int64`|
|name|`utf8`|
|non_overridable_arguments|`json`|
|notification_property|`json`|
|number_of_workers|`int64`|
|role|`utf8`|
|security_configuration|`utf8`|
|source_control_details|`json`|
|timeout|`int64`|
|worker_type|`utf8`|