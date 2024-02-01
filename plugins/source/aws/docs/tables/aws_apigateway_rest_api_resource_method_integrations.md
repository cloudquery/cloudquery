# Table: aws_apigateway_rest_api_resource_method_integrations

This table shows data for Amazon API Gateway Rest API Resource Method Integrations.

https://docs.aws.amazon.com/apigateway/latest/api/API_Integration.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

This table depends on [aws_apigateway_rest_api_resource_methods](aws_apigateway_rest_api_resource_methods.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|rest_api_arn|`utf8`|
|resource_arn|`utf8`|
|method_arn|`utf8`|
|arn|`utf8`|
|cache_key_parameters|`list<item: utf8, nullable>`|
|cache_namespace|`utf8`|
|connection_id|`utf8`|
|connection_type|`utf8`|
|content_handling|`utf8`|
|credentials|`utf8`|
|http_method|`utf8`|
|integration_responses|`json`|
|passthrough_behavior|`utf8`|
|request_parameters|`json`|
|request_templates|`json`|
|timeout_in_millis|`int64`|
|tls_config|`json`|
|type|`utf8`|
|uri|`utf8`|