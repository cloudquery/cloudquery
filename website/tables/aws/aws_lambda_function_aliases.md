# Table: aws_lambda_function_aliases

This table shows data for AWS Lambda Function Aliases.

https://docs.aws.amazon.com/lambda/latest/dg/API_AliasConfiguration.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_lambda_functions](aws_lambda_functions).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|function_arn|`utf8`|
|arn (PK)|`utf8`|
|alias_arn|`utf8`|
|description|`utf8`|
|function_version|`utf8`|
|name|`utf8`|
|revision_id|`utf8`|
|routing_config|`json`|
|url_config|`json`|