# Table: aws_apigatewayv2_api_authorizers

This table shows data for Amazon API Gateway v2 API Authorizers.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-authorizers.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

This table depends on [aws_apigatewayv2_apis](aws_apigatewayv2_apis.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|api_arn|`utf8`|
|api_id|`utf8`|
|arn|`utf8`|
|name|`utf8`|
|authorizer_credentials_arn|`utf8`|
|authorizer_id|`utf8`|
|authorizer_payload_format_version|`utf8`|
|authorizer_result_ttl_in_seconds|`int64`|
|authorizer_type|`utf8`|
|authorizer_uri|`utf8`|
|enable_simple_responses|`bool`|
|identity_source|`list<item: utf8, nullable>`|
|identity_validation_expression|`utf8`|
|jwt_configuration|`json`|