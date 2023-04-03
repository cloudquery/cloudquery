# Table: aws_sns_subscriptions

This table shows data for Sns Subscriptions.

https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html

The composite primary key for this table is (**endpoint**, **owner**, **protocol**, **topic_arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn|String|
|delivery_policy|JSON|
|effective_delivery_policy|JSON|
|filter_policy|JSON|
|redrive_policy|JSON|
|endpoint (PK)|String|
|owner (PK)|String|
|protocol (PK)|String|
|subscription_arn|String|
|topic_arn (PK)|String|
|confirmation_was_authenticated|Bool|
|pending_confirmation|Bool|
|raw_message_delivery|Bool|
|subscription_role_arn|String|
|unknown_fields|JSON|