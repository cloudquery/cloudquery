# Table: aws_sns_subscriptions

This table shows data for Sns Subscriptions.

https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|arn (PK)|utf8|
|delivery_policy|json|
|effective_delivery_policy|json|
|filter_policy|json|
|redrive_policy|json|
|endpoint|utf8|
|owner|utf8|
|protocol|utf8|
|subscription_arn|utf8|
|topic_arn|utf8|
|confirmation_was_authenticated|bool|
|pending_confirmation|bool|
|raw_message_delivery|bool|
|subscription_role_arn|utf8|
|unknown_fields|json|