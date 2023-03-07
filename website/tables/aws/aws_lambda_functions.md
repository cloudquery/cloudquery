# Table: aws_lambda_functions

https://docs.aws.amazon.com/lambda/latest/dg/API_GetFunction.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_lambda_functions:
  - [aws_lambda_function_aliases](aws_lambda_function_aliases)
  - [aws_lambda_function_concurrency_configs](aws_lambda_function_concurrency_configs)
  - [aws_lambda_function_event_invoke_configs](aws_lambda_function_event_invoke_configs)
  - [aws_lambda_function_event_source_mappings](aws_lambda_function_event_source_mappings)
  - [aws_lambda_function_versions](aws_lambda_function_versions)

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
|policy_revision_id|String|
|policy_document|JSON|
|code_signing_config|JSON|
|code_repository_type|String|
|update_runtime_on|String|
|runtime_version_arn|String|
|code|JSON|
|concurrency|JSON|
|configuration|JSON|
|tags|JSON|
|result_metadata|JSON|