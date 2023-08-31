# Table: aws_shield_subscriptions

This table shows data for Shield Subscriptions.

https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_Subscription.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|subscription_limits|`json`|
|auto_renew|`utf8`|
|end_time|`timestamp[us, tz=UTC]`|
|limits|`json`|
|proactive_engagement_status|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|subscription_arn|`utf8`|
|time_commitment_in_seconds|`int64`|