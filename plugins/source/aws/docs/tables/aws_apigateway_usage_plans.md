# Table: aws_apigateway_usage_plans

This table shows data for Amazon API Gateway Usage Plans.

https://docs.aws.amazon.com/apigateway/latest/api/API_UsagePlan.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

The following tables depend on aws_apigateway_usage_plans:
  - [aws_apigateway_usage_plan_keys](aws_apigateway_usage_plan_keys.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|api_stages|`json`|
|description|`utf8`|
|id|`utf8`|
|name|`utf8`|
|product_code|`utf8`|
|quota|`json`|
|tags|`json`|
|throttle|`json`|