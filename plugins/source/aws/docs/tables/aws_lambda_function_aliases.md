# Table: aws_lambda_function_aliases

This table shows data for AWS Lambda Function Aliases.

https://docs.aws.amazon.com/lambda/latest/dg/API_AliasConfiguration.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

This table depends on [aws_lambda_functions](aws_lambda_functions.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|function_arn|`utf8`|
|arn|`utf8`|
|alias_arn|`utf8`|
|description|`utf8`|
|function_version|`utf8`|
|name|`utf8`|
|revision_id|`utf8`|
|routing_config|`json`|