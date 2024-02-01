# Table: aws_apigatewayv2_api_models

This table shows data for Amazon API Gateway v2 API Models.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-models.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

This table depends on [aws_apigatewayv2_apis](aws_apigatewayv2_apis.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|api_arn|`utf8`|
|api_id|`utf8`|
|arn|`utf8`|
|model_template|`utf8`|
|name|`utf8`|
|content_type|`utf8`|
|description|`utf8`|
|model_id|`utf8`|
|schema|`utf8`|