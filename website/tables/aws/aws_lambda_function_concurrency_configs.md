# Table: aws_lambda_function_concurrency_configs

This table shows data for AWS Lambda Function Concurrency Configs.

https://docs.aws.amazon.com/lambda/latest/dg/API_ProvisionedConcurrencyConfigListItem.html

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
|allocated_provisioned_concurrent_executions|`int64`|
|available_provisioned_concurrent_executions|`int64`|
|last_modified|`utf8`|
|requested_provisioned_concurrent_executions|`int64`|
|status|`utf8`|
|status_reason|`utf8`|