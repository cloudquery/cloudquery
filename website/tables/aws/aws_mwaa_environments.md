# Table: aws_mwaa_environments

This table shows data for Amazon MWAA Environments.

https://docs.aws.amazon.com/mwaa/latest/API/API_Environment.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|airflow_configuration_options|`json`|
|airflow_version|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|dag_s3_path|`utf8`|
|environment_class|`utf8`|
|execution_role_arn|`utf8`|
|kms_key|`utf8`|
|last_update|`json`|
|logging_configuration|`json`|
|max_workers|`int64`|
|min_workers|`int64`|
|name|`utf8`|
|network_configuration|`json`|
|plugins_s3_object_version|`utf8`|
|plugins_s3_path|`utf8`|
|requirements_s3_object_version|`utf8`|
|requirements_s3_path|`utf8`|
|schedulers|`int64`|
|service_role_arn|`utf8`|
|source_bucket_arn|`utf8`|
|startup_script_s3_object_version|`utf8`|
|startup_script_s3_path|`utf8`|
|status|`utf8`|
|tags|`json`|
|webserver_access_mode|`utf8`|
|webserver_url|`utf8`|
|weekly_maintenance_window_start|`utf8`|