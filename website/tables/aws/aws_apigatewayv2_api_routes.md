# Table: aws_apigatewayv2_api_routes

This table shows data for Amazon API Gateway v2 API Routes.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-routes.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigatewayv2_apis](aws_apigatewayv2_apis).

The following tables depend on aws_apigatewayv2_api_routes:
  - [aws_apigatewayv2_api_route_responses](aws_apigatewayv2_api_route_responses)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|api_arn|`utf8`|
|api_id|`utf8`|
|arn (PK)|`utf8`|
|route_key|`utf8`|
|api_gateway_managed|`bool`|
|api_key_required|`bool`|
|authorization_scopes|`list<item: utf8, nullable>`|
|authorization_type|`utf8`|
|authorizer_id|`utf8`|
|model_selection_expression|`utf8`|
|operation_name|`utf8`|
|request_models|`json`|
|request_parameters|`json`|
|route_id|`utf8`|
|route_response_selection_expression|`utf8`|
|target|`utf8`|