# Table: aws_apigateway_rest_api_gateway_responses

This table shows data for Amazon API Gateway Rest API Gateway Responses.

https://docs.aws.amazon.com/apigateway/latest/api/API_GatewayResponse.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

This table depends on [aws_apigateway_rest_apis](aws_apigateway_rest_apis.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|rest_api_arn|`utf8`|
|arn|`utf8`|
|default_response|`bool`|
|response_parameters|`json`|
|response_templates|`json`|
|response_type|`utf8`|
|status_code|`utf8`|