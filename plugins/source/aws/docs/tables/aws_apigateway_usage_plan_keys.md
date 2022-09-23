# Table: aws_apigateway_usage_plan_keys


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_apigateway_usage_plans`](aws_apigateway_usage_plans.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|usage_plan_arn|String|
|arn|String|
|id|String|
|name|String|
|type|String|
|value|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|