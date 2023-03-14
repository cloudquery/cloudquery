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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|arn (PK)|String|
|id|String|
|name|String|
|protocol_type|String|
|route_selection_expression|String|
|api_endpoint|String|
|api_gateway_managed|Bool|
|api_id|String|
|api_key_selection_expression|String|
|cors_configuration|JSON|
|created_date|Timestamp|
|description|String|
|disable_execute_api_endpoint|Bool|
|disable_schema_validation|Bool|
|import_info|StringArray|
|tags|JSON|
|version|String|
|warnings|StringArray|