
# Table: aws_apigatewayv2_apis
Represents an API
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|name|text|The name of the API|
|protocol_type|text|The API protocol|
|route_selection_expression|text|The route selection expression for the API|
|api_endpoint|text|The URI of the API, of the form {api-id}.execute-api.{region}.amazonaws.com|
|api_gateway_managed|boolean|Specifies whether an API is managed by API Gateway|
|id|text|The API ID|
|api_key_selection_expression|text|An API key selection expression|
|cors_configuration_allow_credentials|boolean|Specifies whether credentials are included in the CORS request|
|cors_configuration_allow_headers|text[]|Represents a collection of allowed headers|
|cors_configuration_allow_methods|text[]|Represents a collection of allowed HTTP methods|
|cors_configuration_allow_origins|text[]|Represents a collection of allowed origins|
|cors_configuration_expose_headers|text[]|Represents a collection of exposed headers|
|cors_configuration_max_age|bigint|The number of seconds that the browser should cache preflight request results|
|created_date|timestamp without time zone|The timestamp when the API was created|
|description|text|The description of the API|
|disable_execute_api_endpoint|boolean|Specifies whether clients can invoke your API by using the default execute-api endpoint|
|disable_schema_validation|boolean|Avoid validating models when creating a deployment|
|import_info|text[]|The validation information during API import|
|tags|jsonb|A collection of tags associated with the API|
|version|text|A version identifier for the API|
|warnings|text[]|The warning messages reported when failonwarnings is turned on during API import|
