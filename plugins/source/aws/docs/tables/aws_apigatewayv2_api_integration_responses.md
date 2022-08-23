
# Table: aws_apigatewayv2_api_integration_responses
Represents an integration response
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_integration_cq_id|uuid|Unique CloudQuery ID of aws_apigatewayv2_api_integrations table (FK)|
|integration_id|text|Represents the identifier of an integration|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|integration_response_key|text|The integration response key|
|content_handling_strategy|text|Supported only for WebSocket APIs|
|integration_response_id|text|The integration response ID|
|response_parameters|jsonb|A key-value map specifying response parameters that are passed to the method response from the backend|
|response_templates|jsonb|The collection of response templates for the integration response as a string-to-string map of key-value pairs|
|template_selection_expression|text|The template selection expressions for the integration response|
