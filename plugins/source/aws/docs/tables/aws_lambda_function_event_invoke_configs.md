# Table: aws_lambda_function_event_invoke_configs

https://docs.aws.amazon.com/lambda/latest/dg/API_FunctionEventInvokeConfig.html

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
|destination_config|JSON|
|last_modified|Timestamp|
|maximum_event_age_in_seconds|Int|
|maximum_retry_attempts|Int|