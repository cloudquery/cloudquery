# Table: aws_servicequotas_quotas



The primary key for this table is **arn**.

## Relations
This table depends on [aws_servicequotas_services](aws_servicequotas_services.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|adjustable|Bool|
|error_reason|JSON|
|global_quota|Bool|
|period|JSON|
|quota_code|String|
|quota_name|String|
|service_code|String|
|service_name|String|
|unit|String|
|usage_metric|JSON|
|value|Float|