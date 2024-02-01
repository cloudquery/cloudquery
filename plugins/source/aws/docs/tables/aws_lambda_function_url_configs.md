# Table: aws_lambda_function_url_configs

This table shows data for AWS Lambda Function URL Configs.

https://docs.aws.amazon.com/lambda/latest/dg/API_FunctionUrlConfig.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **function_arn**.
## Relations

This table depends on [aws_lambda_functions](aws_lambda_functions.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|auth_type|`utf8`|
|creation_time|`utf8`|
|function_arn|`utf8`|
|function_url|`utf8`|
|last_modified_time|`utf8`|
|cors|`json`|
|invoke_mode|`utf8`|