# Table: aws_apigatewayv2_api_routes

https://docs.aws.amazon.com/apigateway/latest/api/API_Route.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_apigatewayv2_apis](aws_apigatewayv2_apis.md).

The following tables depend on aws_apigatewayv2_api_routes:
  - [aws_apigatewayv2_api_route_responses](aws_apigatewayv2_api_route_responses.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|api_arn|String|
|api_id|String|
|arn|String|
|route_key|String|
|api_gateway_managed|Bool|
|api_key_required|Bool|
|authorization_scopes|StringArray|
|authorization_type|String|
|authorizer_id|String|
|model_selection_expression|String|
|operation_name|String|
|request_models|JSON|
|request_parameters|JSON|
|route_id|String|
|route_response_selection_expression|String|
|target|String|