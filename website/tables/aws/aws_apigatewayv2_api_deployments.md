# Table: aws_apigatewayv2_api_deployments

This table shows data for Amazon API Gateway v2 API Deployments.

https://docs.aws.amazon.com/apigateway/latest/api/API_Deployment.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigatewayv2_apis](aws_apigatewayv2_apis).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|api_arn|`utf8`|
|api_id|`utf8`|
|arn (PK)|`utf8`|
|auto_deployed|`bool`|
|created_date|`timestamp[us, tz=UTC]`|
|deployment_id|`utf8`|
|deployment_status|`utf8`|
|deployment_status_message|`utf8`|
|description|`utf8`|