# Table: aws_lambda_functions

This table shows data for AWS Lambda Functions.

https://docs.aws.amazon.com/lambda/latest/dg/API_GetFunction.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_lambda_functions:
  - [aws_lambda_function_aliases](aws_lambda_function_aliases.md)
  - [aws_lambda_function_concurrency_configs](aws_lambda_function_concurrency_configs.md)
  - [aws_lambda_function_event_invoke_configs](aws_lambda_function_event_invoke_configs.md)
  - [aws_lambda_function_event_source_mappings](aws_lambda_function_event_source_mappings.md)
  - [aws_lambda_function_url_configs](aws_lambda_function_url_configs.md)
  - [aws_lambda_function_versions](aws_lambda_function_versions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|policy_revision_id|`utf8`|
|policy_document|`json`|
|code_signing_config|`json`|
|code_repository_type|`utf8`|
|update_runtime_on|`utf8`|
|runtime_version_arn|`utf8`|
|code|`json`|
|concurrency|`json`|
|configuration|`json`|
|tags|`json`|