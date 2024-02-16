# Table: aws_apigateway_rest_api_resources

This table shows data for Amazon API Gateway Rest API Resources.

https://docs.aws.amazon.com/apigateway/latest/api/API_Resource.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

This table depends on [aws_apigateway_rest_apis](aws_apigateway_rest_apis.md).

The following tables depend on aws_apigateway_rest_api_resources:
  - [aws_apigateway_rest_api_resource_methods](aws_apigateway_rest_api_resource_methods.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|rest_api_arn|`utf8`|
|arn|`utf8`|
|id|`utf8`|
|parent_id|`utf8`|
|path|`utf8`|
|path_part|`utf8`|
|resource_methods|`json`|