# Table: aws_apigatewayv2_apis

This table shows data for Amazon API Gateway v2 APIs.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

The following tables depend on aws_apigatewayv2_apis:
  - [aws_apigatewayv2_api_authorizers](aws_apigatewayv2_api_authorizers.md)
  - [aws_apigatewayv2_api_deployments](aws_apigatewayv2_api_deployments.md)
  - [aws_apigatewayv2_api_integrations](aws_apigatewayv2_api_integrations.md)
  - [aws_apigatewayv2_api_models](aws_apigatewayv2_api_models.md)
  - [aws_apigatewayv2_api_routes](aws_apigatewayv2_api_routes.md)
  - [aws_apigatewayv2_api_stages](aws_apigatewayv2_api_stages.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
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