# Table: aws_apigateway_rest_api_resource_methods

This table shows data for Amazon API Gateway Rest API Resource Methods.

https://docs.aws.amazon.com/apigateway/latest/api/API_Method.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigateway_rest_api_resources](aws_apigateway_rest_api_resources).

The following tables depend on aws_apigateway_rest_api_resource_methods:
  - [aws_apigateway_rest_api_resource_method_integrations](aws_apigateway_rest_api_resource_method_integrations)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|rest_api_arn|`utf8`|
|resource_arn|`utf8`|
|arn (PK)|`utf8`|
|api_key_required|`bool`|
|authorization_scopes|`list<item: utf8, nullable>`|
|authorization_type|`utf8`|
|authorizer_id|`utf8`|
|http_method|`utf8`|
|method_integration|`json`|
|method_responses|`json`|
|operation_name|`utf8`|
|request_models|`json`|
|request_parameters|`json`|
|request_validator_id|`utf8`|