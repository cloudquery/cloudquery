# Table: aws_apigateway_usage_plan_keys

https://docs.aws.amazon.com/apigateway/latest/api/API_UsagePlanKey.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigateway_usage_plans](aws_apigateway_usage_plans.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|usage_plan_arn|String|
|arn (PK)|String|
|id|String|
|name|String|
|type|String|
|value|String|