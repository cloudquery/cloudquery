# Table: aws_apigateway_rest_api_resources

This table shows data for Amazon API Gateway Rest API Resources.

https://docs.aws.amazon.com/apigateway/latest/api/API_Resource.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigateway_rest_apis](aws_apigateway_rest_apis).

The following tables depend on aws_apigateway_rest_api_resources:
  - [aws_apigateway_rest_api_resource_methods](aws_apigateway_rest_api_resource_methods)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|rest_api_arn|`utf8`|
|arn (PK)|`utf8`|
|id|`utf8`|
|parent_id|`utf8`|
|path|`utf8`|
|path_part|`utf8`|
|resource_methods|`json`|