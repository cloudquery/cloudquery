# Table: aws_apigateway_rest_api_authorizers

This table shows data for Amazon API Gateway Rest API Authorizers.

https://docs.aws.amazon.com/apigateway/latest/api/API_Authorizer.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigateway_rest_apis](aws_apigateway_rest_apis).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|rest_api_arn|`utf8`|
|arn (PK)|`utf8`|
|auth_type|`utf8`|
|authorizer_credentials|`utf8`|
|authorizer_result_ttl_in_seconds|`int64`|
|authorizer_uri|`utf8`|
|id|`utf8`|
|identity_source|`utf8`|
|identity_validation_expression|`utf8`|
|name|`utf8`|
|provider_arns|`list<item: utf8, nullable>`|
|type|`utf8`|