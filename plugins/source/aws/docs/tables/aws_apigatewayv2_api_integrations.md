
# Table: aws_apigatewayv2_api_integrations
Represents an integration
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_cq_id|uuid|Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)|
|api_id|text|The API id|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|api_gateway_managed|boolean|Specifies whether an integration is managed by API Gateway|
|connection_id|text|The ID of the VPC link for a private integration|
|connection_type|text|The type of the network connection to the integration endpoint|
|content_handling_strategy|text|Supported only for WebSocket APIs|
|credentials_arn|text|Specifies the credentials required for the integration, if any|
|description|text|Represents the description of an integration|
|integration_id|text|Represents the identifier of an integration|
|integration_method|text|Specifies the integration's HTTP method type|
|integration_response_selection_expression|text|The integration response selection expression for the integration|
|integration_subtype|text|Supported only for HTTP API AWS_PROXY integrations|
|integration_type|text|The integration type of an integration|
|integration_uri|text|For a Lambda integration, specify the URI of a Lambda function|
|passthrough_behavior|text|Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the requestTemplates property on the Integration resource|
|payload_format_version|text|Specifies the format of the payload sent to an integration|
|request_parameters|jsonb|For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend|
|request_templates|jsonb|Represents a map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client|
|response_parameters|jsonb|Supported only for HTTP APIs|
|template_selection_expression|text|The template selection expression for the integration|
|timeout_in_millis|bigint|Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs|
|tls_config_server_name_to_verify|text|If you specify a server name, API Gateway uses it to verify the hostname on the integration's certificate|
