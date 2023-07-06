# Table: aws_apigatewayv2_api_models

This table shows data for Amazon API Gateway v2 API Models.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-models.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigatewayv2_apis](aws_apigatewayv2_apis).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|api_arn|`utf8`|
|api_id|`utf8`|
|arn (PK)|`utf8`|
|model_template|`utf8`|
|name|`utf8`|
|content_type|`utf8`|
|description|`utf8`|
|model_id|`utf8`|
|schema|`utf8`|