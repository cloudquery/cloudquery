# Table: aws_apigateway_rest_api_models

This table shows data for Amazon API Gateway Rest API Models.

https://docs.aws.amazon.com/apigateway/latest/api/API_Model.html

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
|model_template|`utf8`|
|content_type|`utf8`|
|description|`utf8`|
|id|`utf8`|
|name|`utf8`|
|schema|`utf8`|