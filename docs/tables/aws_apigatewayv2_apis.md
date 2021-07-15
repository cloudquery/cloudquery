
# Table: aws_apigatewayv2_apis
Represents an API.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|name|text|The name of the API.|
|protocol_type|text|The API protocol.|
|route_selection_expression|text|The route selection expression for the API. For HTTP APIs, the routeSelectionExpression must be ${request.method} ${request.path}. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.|
|api_endpoint|text|The URI of the API, of the form {api-id}.execute-api.{region}.amazonaws.com. The stage name is typically appended to this URI to form a complete path to a deployed API stage.|
|api_gateway_managed|boolean|Specifies whether an API is managed by API Gateway. You can't update or delete a managed API by using API Gateway. A managed API can be deleted only through the tooling or service that created it.|
|id|text|The API ID.|
|api_key_selection_expression|text|An API key selection expression. Supported only for WebSocket APIs. See API Key Selection Expressions (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).|
|cors_configuration_allow_credentials|boolean|Specifies whether credentials are included in the CORS request. Supported only for HTTP APIs.|
|cors_configuration_allow_headers|text[]|Represents a collection of allowed headers. Supported only for HTTP APIs.|
|cors_configuration_allow_methods|text[]|Represents a collection of allowed HTTP methods. Supported only for HTTP APIs.|
|cors_configuration_allow_origins|text[]|Represents a collection of allowed origins. Supported only for HTTP APIs.|
|cors_configuration_expose_headers|text[]|Represents a collection of exposed headers. Supported only for HTTP APIs.|
|cors_configuration_max_age|integer|The number of seconds that the browser should cache preflight request results. Supported only for HTTP APIs.|
|created_date|timestamp without time zone|The timestamp when the API was created.|
|description|text|The description of the API.|
|disable_execute_api_endpoint|boolean|Specifies whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.|
|disable_schema_validation|boolean|Avoid validating models when creating a deployment. Supported only for WebSocket APIs.|
|import_info|text[]|The validation information during API import. This may include particular properties of your OpenAPI definition which are ignored during import. Supported only for HTTP APIs.|
|tags|jsonb|A collection of tags associated with the API.|
|version|text|A version identifier for the API.|
|warnings|text[]|The warning messages reported when failonwarnings is turned on during API import.|
