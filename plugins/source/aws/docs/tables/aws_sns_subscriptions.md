# Table: aws_sns_subscriptions



The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|arn (PK)|String|
|endpoint|String|
|owner|String|
|protocol|String|
|topic_arn|String|
|confirmation_was_authenticated|Bool|
|delivery_policy|String|
|effective_delivery_policy|String|
|filter_policy|String|
|pending_confirmation|Bool|
|raw_message_delivery|Bool|
|redrive_policy|String|
|subscription_role_arn|String|
|unknown_fields|JSON|