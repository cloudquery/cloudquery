# Table: aws_lambda_functions



The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_lambda_functions:
  - [aws_lambda_function_event_invoke_configs](aws_lambda_function_event_invoke_configs.md)
  - [aws_lambda_function_aliases](aws_lambda_function_aliases.md)
  - [aws_lambda_function_versions](aws_lambda_function_versions.md)
  - [aws_lambda_function_concurrency_configs](aws_lambda_function_concurrency_configs.md)
  - [aws_lambda_function_event_source_mappings](aws_lambda_function_event_source_mappings.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn|String|
|policy_revision_id|String|
|policy_document|JSON|
|code_signing_config|JSON|
|code_repository_type|String|
|code|JSON|
|concurrency|JSON|
|configuration|JSON|
|tags|JSON|
|result_metadata|JSON|