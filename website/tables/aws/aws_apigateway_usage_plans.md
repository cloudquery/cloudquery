# Table: aws_apigateway_usage_plans

This table shows data for Amazon API Gateway Usage Plans.

https://docs.aws.amazon.com/apigateway/latest/api/API_UsagePlan.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

The following tables depend on aws_apigateway_usage_plans:
  - [aws_apigateway_usage_plan_keys](aws_apigateway_usage_plan_keys)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|api_stages|`json`|
|description|`utf8`|
|id|`utf8`|
|name|`utf8`|
|product_code|`utf8`|
|quota|`json`|
|tags|`json`|
|throttle|`json`|