
# Table: aws_apigatewayv2_api_authorizers
Represents an authorizer
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_cq_id|uuid|Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)|
|api_id|text|The API id|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|name|text|The name of the authorizer|
|authorizer_credentials_arn|text|Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer|
|authorizer_id|text|The authorizer identifier|
|authorizer_payload_format_version|text|Specifies the format of the payload sent to an HTTP API Lambda authorizer|
|authorizer_result_ttl_in_seconds|bigint|The time to live (TTL) for cached authorizer results, in seconds|
|authorizer_type|text|The authorizer type|
|authorizer_uri|text|The authorizer's Uniform Resource Identifier (URI)|
|enable_simple_responses|boolean|Specifies whether a Lambda authorizer returns a response in a simple format|
|identity_source|text[]|The identity source for which authorization is requested|
|identity_validation_expression|text|The validation expression does not apply to the REQUEST authorizer|
|jwt_configuration_audience|text[]|A list of the intended recipients of the JWT|
|jwt_configuration_issuer|text|The base domain of the identity provider that issues JSON Web Tokens|
