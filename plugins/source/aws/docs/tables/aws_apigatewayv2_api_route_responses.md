
# Table: aws_apigatewayv2_api_route_responses
Represents a route response.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_route_cq_id|uuid|Unique CloudQuery ID of aws_apigatewayv2_api_routes table (FK)|
|route_id|text|Represents the identifier of an route.|
|arn|text|The Amazon Resource Name (ARN) for the resource.|
|route_response_key|text|Represents the route response key of a route response.|
|model_selection_expression|text|Represents the model selection expression of a route response. Supported only for WebSocket APIs.|
|response_models|jsonb|Represents the response models of a route response.|
|response_parameters|jsonb|Represents the response parameters of a route response.|
|route_response_id|text|Represents the identifier of a route response.|
