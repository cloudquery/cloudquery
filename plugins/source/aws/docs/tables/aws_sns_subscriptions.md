# Table: aws_sns_subscriptions



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
|delivery_policy|JSON|
|effective_delivery_policy|JSON|
|filter_policy|JSON|
|redrive_policy|JSON|
|endpoint|String|
|owner|String|
|protocol|String|
|topic_arn|String|
|confirmation_was_authenticated|Bool|
|pending_confirmation|Bool|
|raw_message_delivery|Bool|
|subscription_role_arn|String|
|unknown_fields|JSON|