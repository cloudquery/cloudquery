# Table: aws_lambda_function_url_configs

This table shows data for AWS Lambda Function URL Configs.

https://docs.aws.amazon.com/lambda/latest/dg/API_FunctionUrlConfig.html

The composite primary key for this table is (**function_arn**, **function_url**).

## Relations

This table depends on [aws_lambda_functions](aws_lambda_functions.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|auth_type|`utf8`|
|creation_time|`utf8`|
|function_arn (PK)|`utf8`|
|function_url (PK)|`utf8`|
|last_modified_time|`utf8`|
|cors|`json`|
|invoke_mode|`utf8`|