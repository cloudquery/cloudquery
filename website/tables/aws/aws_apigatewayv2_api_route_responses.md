# Table: aws_apigatewayv2_api_route_responses

This table shows data for AWS API Gateway v2 API Route Responses.

https://docs.aws.amazon.com/apigateway/latest/api/API_RouteResponse.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigatewayv2_api_routes](aws_apigatewayv2_api_routes).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|api_route_arn|String|
|route_id|String|
|arn (PK)|String|
|route_response_key|String|
|model_selection_expression|String|
|response_models|JSON|
|response_parameters|JSON|
|route_response_id|String|