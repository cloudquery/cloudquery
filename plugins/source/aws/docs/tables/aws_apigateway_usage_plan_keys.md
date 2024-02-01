# Table: aws_apigateway_usage_plan_keys

This table shows data for Amazon API Gateway Usage Plan Keys.

https://docs.aws.amazon.com/apigateway/latest/api/API_UsagePlanKey.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

This table depends on [aws_apigateway_usage_plans](aws_apigateway_usage_plans.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|usage_plan_arn|`utf8`|
|arn|`utf8`|
|id|`utf8`|
|name|`utf8`|
|type|`utf8`|
|value|`utf8`|