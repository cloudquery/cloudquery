# Table: aws_apigateway_rest_api_gateway_responses

This table shows data for Amazon API Gateway Rest API Gateway Responses.

https://docs.aws.amazon.com/apigateway/latest/api/API_GatewayResponse.html

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
|default_response|`bool`|
|response_parameters|`json`|
|response_templates|`json`|
|response_type|`utf8`|
|status_code|`utf8`|