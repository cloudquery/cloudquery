# Table: aws_apigateway_usage_plan_keys

This table shows data for Amazon API Gateway Usage Plan Keys.

https://docs.aws.amazon.com/apigateway/latest/api/API_UsagePlanKey.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigateway_usage_plans](aws_apigateway_usage_plans).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|usage_plan_arn|`utf8`|
|arn (PK)|`utf8`|
|id|`utf8`|
|name|`utf8`|
|type|`utf8`|
|value|`utf8`|