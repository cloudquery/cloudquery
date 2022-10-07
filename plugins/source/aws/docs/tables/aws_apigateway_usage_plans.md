# Table: aws_apigateway_usage_plans



The primary key for this table is **_cq_id**.

## Relations
The following tables depend on `aws_apigateway_usage_plans`:
  - [`aws_apigateway_usage_plan_keys`](aws_apigateway_usage_plan_keys.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|arn|String|
|api_stages|JSON|
|description|String|
|id|String|
|name|String|
|product_code|String|
|quota|JSON|
|tags|JSON|
|throttle|JSON|