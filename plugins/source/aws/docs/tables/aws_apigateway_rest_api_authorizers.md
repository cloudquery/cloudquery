
# Table: aws_apigateway_rest_api_authorizers
Represents an authorization layer for methods
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_cq_id|uuid|Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)|
|rest_api_id|text|The API's identifier|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|auth_type|text|Optional customer-defined field, used in OpenAPI imports and exports without functional impact|
|authorizer_credentials|text|Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer|
|authorizer_result_ttl_in_seconds|bigint|The TTL in seconds of cached authorizer results|
|authorizer_uri|text|Specifies the authorizer's Uniform Resource Identifier (URI)|
|id|text|The identifier for the authorizer resource|
|identity_source|text|The identity source for which authorization is requested|
|identity_validation_expression|text|A validation expression for the incoming identity token|
|name|text|The name of the authorizer|
|provider_arns|text[]|A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS authorizer|
|type|text|The authorizer type|
