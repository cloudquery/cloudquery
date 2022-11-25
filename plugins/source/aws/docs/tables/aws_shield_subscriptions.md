# Table: aws_shield_subscriptions

https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_Subscription.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|subscription_limits|JSON|
|auto_renew|String|
|end_time|Timestamp|
|limits|JSON|
|proactive_engagement_status|String|
|start_time|Timestamp|
|time_commitment_in_seconds|Int|