# Table: aws_sns_topics



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
|tags|JSON|
|delivery_policy|JSON|
|policy|JSON|
|effective_delivery_policy|JSON|
|display_name|String|
|owner|String|
|subscriptions_confirmed|Int|
|subscriptions_deleted|Int|
|subscriptions_pending|Int|
|kms_master_key_id|String|
|fifo_topic|Bool|
|content_based_deduplication|Bool|
|unknown_fields|JSON|