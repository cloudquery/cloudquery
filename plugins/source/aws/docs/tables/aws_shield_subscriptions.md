# Table: aws_shield_subscriptions



The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|arn (PK)|String|
|subscription_limits|JSON|
|auto_renew|String|
|end_time|Timestamp|
|limits|JSON|
|proactive_engagement_status|String|
|start_time|Timestamp|
|time_commitment_in_seconds|Int|