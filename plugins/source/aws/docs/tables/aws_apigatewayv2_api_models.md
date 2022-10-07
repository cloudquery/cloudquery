# Table: aws_apigatewayv2_api_models



The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_apigatewayv2_apis`](aws_apigatewayv2_apis.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|api_arn|String|
|api_id|String|
|arn|String|
|model_template|String|
|name|String|
|content_type|String|
|description|String|
|model_id|String|
|schema|String|