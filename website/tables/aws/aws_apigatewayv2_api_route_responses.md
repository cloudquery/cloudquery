# Table: aws_apigatewayv2_api_route_responses

This table shows data for Amazon API Gateway v2 API Route Responses.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-routes-routeid-routeresponses.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigatewayv2_api_routes](aws_apigatewayv2_api_routes).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|api_route_arn|`utf8`|
|route_id|`utf8`|
|arn (PK)|`utf8`|
|route_response_key|`utf8`|
|model_selection_expression|`utf8`|
|response_models|`json`|
|response_parameters|`json`|
|route_response_id|`utf8`|