# Table: aws_apigateway_rest_api_documentation_versions

This table shows data for Amazon API Gateway Rest API Documentation Versions.

https://docs.aws.amazon.com/apigateway/latest/api/API_DocumentationVersion.html

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
|created_date|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|version|`utf8`|