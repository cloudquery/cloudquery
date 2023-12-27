# Table: aws_lambda_function_event_invoke_configs

This table shows data for AWS Lambda Function Event Invoke Configs.

https://docs.aws.amazon.com/lambda/latest/dg/API_FunctionEventInvokeConfig.html

The primary key for this table is **function_arn**.

## Relations

This table depends on [aws_lambda_functions](aws_lambda_functions.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|destination_config|`json`|
|function_arn (PK)|`utf8`|
|last_modified|`timestamp[us, tz=UTC]`|
|maximum_event_age_in_seconds|`int64`|
|maximum_retry_attempts|`int64`|