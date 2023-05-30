# Table: aws_apigatewayv2_apis

This table shows data for Amazon API Gateway v2 APIs.

https://docs.aws.amazon.com/apigateway/latest/api/API_Api.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

The following tables depend on aws_apigatewayv2_apis:
  - [aws_apigatewayv2_api_authorizers](aws_apigatewayv2_api_authorizers)
  - [aws_apigatewayv2_api_deployments](aws_apigatewayv2_api_deployments)
  - [aws_apigatewayv2_api_integrations](aws_apigatewayv2_api_integrations)
  - [aws_apigatewayv2_api_models](aws_apigatewayv2_api_models)
  - [aws_apigatewayv2_api_routes](aws_apigatewayv2_api_routes)
  - [aws_apigatewayv2_api_stages](aws_apigatewayv2_api_stages)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|id|`utf8`|
|name|`utf8`|
|protocol_type|`utf8`|
|route_selection_expression|`utf8`|
|api_endpoint|`utf8`|
|api_gateway_managed|`bool`|
|api_id|`utf8`|
|api_key_selection_expression|`utf8`|
|cors_configuration|`json`|
|created_date|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|disable_execute_api_endpoint|`bool`|
|disable_schema_validation|`bool`|
|import_info|`list<item: utf8, nullable>`|
|tags|`json`|
|version|`utf8`|
|warnings|`list<item: utf8, nullable>`|