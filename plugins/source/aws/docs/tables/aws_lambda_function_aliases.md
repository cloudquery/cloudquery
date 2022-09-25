# Table: aws_lambda_function_aliases


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_lambda_functions`](aws_lambda_functions.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|function_arn|String|
|alias_arn|String|
|description|String|
|function_version|String|
|name|String|
|revision_id|String|
|routing_config|JSON|
|url_config|JSON|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|