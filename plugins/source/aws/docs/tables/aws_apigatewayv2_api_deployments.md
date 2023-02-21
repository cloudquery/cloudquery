# Table: aws_apigatewayv2_api_deployments

https://docs.aws.amazon.com/apigateway/latest/api/API_Deployment.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigatewayv2_apis](aws_apigatewayv2_apis.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|api_arn|String|
|api_id|String|
|arn (PK)|String|
|auto_deployed|Bool|
|created_date|Timestamp|
|deployment_id|String|
|deployment_status|String|
|deployment_status_message|String|
|description|String|