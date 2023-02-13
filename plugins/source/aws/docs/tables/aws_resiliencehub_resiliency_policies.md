# Table: aws_resiliencehub_resiliency_policies

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_ResiliencyPolicy.html

The primary key for this table is **arn**.

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
|creation_time|Timestamp|
|data_location_constraint|String|
|estimated_cost_tier|String|
|policy|JSON|
|policy_arn|String|
|policy_description|String|
|policy_name|String|
|tags|JSON|
|tier|String|