# Table: aws_apigateway_rest_apis

This table shows data for Amazon API Gateway Rest APIs.

https://docs.aws.amazon.com/apigateway/latest/api/API_RestApi.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_apigateway_rest_apis:
  - [aws_apigateway_rest_api_authorizers](aws_apigateway_rest_api_authorizers.md)
  - [aws_apigateway_rest_api_deployments](aws_apigateway_rest_api_deployments.md)
  - [aws_apigateway_rest_api_documentation_parts](aws_apigateway_rest_api_documentation_parts.md)
  - [aws_apigateway_rest_api_documentation_versions](aws_apigateway_rest_api_documentation_versions.md)
  - [aws_apigateway_rest_api_gateway_responses](aws_apigateway_rest_api_gateway_responses.md)
  - [aws_apigateway_rest_api_models](aws_apigateway_rest_api_models.md)
  - [aws_apigateway_rest_api_request_validators](aws_apigateway_rest_api_request_validators.md)
  - [aws_apigateway_rest_api_resources](aws_apigateway_rest_api_resources.md)
  - [aws_apigateway_rest_api_stages](aws_apigateway_rest_api_stages.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|api_key_source|`utf8`|
|binary_media_types|`list<item: utf8, nullable>`|
|created_date|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|disable_execute_api_endpoint|`bool`|
|endpoint_configuration|`json`|
|id|`utf8`|
|minimum_compression_size|`int64`|
|name|`utf8`|
|policy|`utf8`|
|root_resource_id|`utf8`|
|tags|`json`|
|version|`utf8`|
|warnings|`list<item: utf8, nullable>`|