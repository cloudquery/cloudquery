# Table: aws_mwaa_environments

https://docs.aws.amazon.com/mwaa/latest/API/API_Environment.html

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
|arn (PK)|String|
|airflow_configuration_options|JSON|
|airflow_version|String|
|created_at|Timestamp|
|dag_s3_path|String|
|environment_class|String|
|execution_role_arn|String|
|kms_key|String|
|last_update|JSON|
|logging_configuration|JSON|
|max_workers|Int|
|min_workers|Int|
|name|String|
|network_configuration|JSON|
|plugins_s3_object_version|String|
|plugins_s3_path|String|
|requirements_s3_object_version|String|
|requirements_s3_path|String|
|schedulers|Int|
|service_role_arn|String|
|source_bucket_arn|String|
|status|String|
|tags|JSON|
|webserver_access_mode|String|
|webserver_url|String|
|weekly_maintenance_window_start|String|