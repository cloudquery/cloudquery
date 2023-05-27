# Table: aws_apigatewayv2_api_stages

This table shows data for Amazon API Gateway v2 API Stages.

https://docs.aws.amazon.com/apigateway/latest/api/API_Stage.html

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
|stage_name|`utf8`|
|access_log_settings|`json`|
|api_gateway_managed|`bool`|
|auto_deploy|`bool`|
|client_certificate_id|`utf8`|
|created_date|`timestamp[us, tz=UTC]`|
|default_route_settings|`json`|
|deployment_id|`utf8`|
|description|`utf8`|
|last_deployment_status_message|`utf8`|
|last_updated_date|`timestamp[us, tz=UTC]`|
|route_settings|`json`|
|stage_variables|`json`|
|tags|`json`|