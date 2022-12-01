# Table: aws_lambda_function_concurrency_configs

https://docs.aws.amazon.com/lambda/latest/dg/API_ProvisionedConcurrencyConfigListItem.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_lambda_functions](aws_lambda_functions.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|function_arn|String|
|allocated_provisioned_concurrent_executions|Int|
|available_provisioned_concurrent_executions|Int|
|last_modified|String|
|requested_provisioned_concurrent_executions|Int|
|status|String|
|status_reason|String|