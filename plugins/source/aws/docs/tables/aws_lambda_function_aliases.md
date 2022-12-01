# Table: aws_lambda_function_aliases



The primary key for this table is **arn**.

## Relations
This table depends on [aws_lambda_functions](aws_lambda_functions.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|function_arn|String|
|arn (PK)|String|
|description|String|
|function_version|String|
|name|String|
|revision_id|String|
|routing_config|JSON|
|url_config|JSON|