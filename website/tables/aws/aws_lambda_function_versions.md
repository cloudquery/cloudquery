# Table: aws_lambda_function_versions

This table shows data for AWS Lambda Function Versions.

https://docs.aws.amazon.com/lambda/latest/dg/API_FunctionConfiguration.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_lambda_functions](aws_lambda_functions).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|function_arn|`utf8`|
|architectures|`list<item: utf8, nullable>`|
|code_sha256|`utf8`|
|code_size|`int64`|
|dead_letter_config|`json`|
|description|`utf8`|
|environment|`json`|
|ephemeral_storage|`json`|
|file_system_configs|`json`|
|function_name|`utf8`|
|handler|`utf8`|
|image_config_response|`json`|
|kms_key_arn|`utf8`|
|last_modified|`utf8`|
|last_update_status|`utf8`|
|last_update_status_reason|`utf8`|
|last_update_status_reason_code|`utf8`|
|layers|`json`|
|master_arn|`utf8`|
|memory_size|`int64`|
|package_type|`utf8`|
|revision_id|`utf8`|
|role|`utf8`|
|runtime|`utf8`|
|runtime_version_config|`json`|
|signing_job_arn|`utf8`|
|signing_profile_version_arn|`utf8`|
|snap_start|`json`|
|state|`utf8`|
|state_reason|`utf8`|
|state_reason_code|`utf8`|
|timeout|`int64`|
|tracing_config|`json`|
|version|`utf8`|
|vpc_config|`json`|