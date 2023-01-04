# Table: aws_apigatewayv2_api_route_responses

https://docs.aws.amazon.com/apigateway/latest/api/API_RouteResponse.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_apigatewayv2_api_routes](aws_apigatewayv2_api_routes.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|api_route_arn|String|
|arn|String|
|route_response_key|String|
|model_selection_expression|String|
|response_models|JSON|
|response_parameters|JSON|
|route_response_id|String|