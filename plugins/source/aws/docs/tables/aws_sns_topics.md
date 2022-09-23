# Table: aws_sns_topics


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|delivery_policy|String|
|display_name|String|
|owner|String|
|policy|String|
|subscriptions_confirmed|Int|
|subscriptions_deleted|Int|
|subscriptions_pending|Int|
|effective_delivery_policy|String|
|kms_master_key_id|String|
|fifo_topic|Bool|
|content_based_deduplication|Bool|
|unknown_fields|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|