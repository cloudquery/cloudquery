
# Table: aws_apigatewayv2_api_integration_responses
Represents an integration response.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_integration_cq_id|uuid|Unique CloudQuery ID of aws_apigatewayv2_api_integrations table (FK)|
|integration_id|text|Represents the identifier of an integration.|
|integration_response_key|text|The integration response key.|
|content_handling_strategy|text|Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are CONVERT_TO_BINARY and CONVERT_TO_TEXT, with the following behaviors: CONVERT_TO_BINARY: Converts a response payload from a Base64-encoded string to the corresponding binary blob. CONVERT_TO_TEXT: Converts a response payload from a binary blob to a Base64-encoded string. If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.|
|integration_response_id|text|The integration response ID.|
|response_parameters|jsonb|A key-value map specifying response parameters that are passed to the method response from the backend. The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of method.response.header.{name}, where name is a valid and unique header name. The mapped non-static value must match the pattern of integration.response.header.{name} or integration.response.body.{JSON-expression}, where name is a valid and unique response header name and JSON-expression is a valid JSON expression without the $ prefix.|
|response_templates|jsonb|The collection of response templates for the integration response as a string-to-string map of key-value pairs. Response templates are represented as a key/value map, with a content-type as the key and a template as the value.|
|template_selection_expression|text|The template selection expressions for the integration response.|
