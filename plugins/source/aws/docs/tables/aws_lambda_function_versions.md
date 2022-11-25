# Table: aws_lambda_function_versions

https://docs.aws.amazon.com/lambda/latest/dg/API_FunctionConfiguration.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_lambda_functions](aws_lambda_functions.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|function_arn|String|
|architectures|StringArray|
|code_sha256|String|
|code_size|Int|
|dead_letter_config|JSON|
|description|String|
|environment|JSON|
|ephemeral_storage|JSON|
|file_system_configs|JSON|
|function_name|String|
|handler|String|
|image_config_response|JSON|
|kms_key_arn|String|
|last_modified|String|
|last_update_status|String|
|last_update_status_reason|String|
|last_update_status_reason_code|String|
|layers|JSON|
|master_arn|String|
|memory_size|Int|
|package_type|String|
|revision_id|String|
|role|String|
|runtime|String|
|signing_job_arn|String|
|signing_profile_version_arn|String|
|state|String|
|state_reason|String|
|state_reason_code|String|
|timeout|Int|
|tracing_config|JSON|
|version|String|
|vpc_config|JSON|