
# Table: aws_apigatewayv2_api_routes
Represents a route
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_cq_id|uuid|Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)|
|api_id|text|The API id|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|route_key|text|The route key for the route|
|api_gateway_managed|boolean|Specifies whether a route is managed by API Gateway|
|api_key_required|boolean|Specifies whether an API key is required for this route|
|authorization_scopes|text[]|A list of authorization scopes configured on a route|
|authorization_type|text|The authorization type for the route|
|authorizer_id|text|The identifier of the Authorizer resource to be associated with this route|
|model_selection_expression|text|The model selection expression for the route|
|operation_name|text|The operation name for the route|
|request_models|jsonb|The request models for the route|
|request_parameters|jsonb|The request parameters for the route|
|route_id|text|The route ID|
|route_response_selection_expression|text|The route response selection expression for the route|
|target|text|The target for the route|
