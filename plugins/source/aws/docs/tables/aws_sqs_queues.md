# Table: aws_sqs_queues



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
|policy|JSON|
|redrive_policy|JSON|
|redrive_allow_policy|JSON|
|url|String|
|approximate_number_of_messages|Int|
|approximate_number_of_messages_delayed|Int|
|approximate_number_of_messages_not_visible|Int|
|created_timestamp|Int|
|delay_seconds|Int|
|last_modified_timestamp|Int|
|maximum_message_size|Int|
|message_retention_period|Int|
|receive_message_wait_time_seconds|Int|
|visibility_timeout|Int|
|kms_master_key_id|String|
|kms_data_key_reuse_period_seconds|Int|
|sqs_managed_sse_enabled|Bool|
|fifo_queue|Bool|
|content_based_deduplication|Bool|
|deduplication_scope|String|
|fifo_throughput_limit|String|
|unknown_fields|JSON|