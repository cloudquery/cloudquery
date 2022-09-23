# Table: aws_apigatewayv2_api_deployments


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_apigatewayv2_apis`](aws_apigatewayv2_apis.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|api_arn|String|
|api_id|String|
|arn|String|
|auto_deployed|Bool|
|created_date|Timestamp|
|deployment_id|String|
|deployment_status|String|
|deployment_status_message|String|
|description|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|