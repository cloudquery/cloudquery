# Table: aws_sqs_queues

This table shows data for Sqs Queues.

https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_GetQueueAttributes.html

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
|policy|json|
|redrive_policy|json|
|redrive_allow_policy|json|
|tags|json|
|url|utf8|
|approximate_number_of_messages|int64|
|approximate_number_of_messages_delayed|int64|
|approximate_number_of_messages_not_visible|int64|
|created_timestamp|int64|
|delay_seconds|int64|
|last_modified_timestamp|int64|
|maximum_message_size|int64|
|message_retention_period|int64|
|receive_message_wait_time_seconds|int64|
|visibility_timeout|int64|
|kms_master_key_id|utf8|
|kms_data_key_reuse_period_seconds|int64|
|sqs_managed_sse_enabled|bool|
|fifo_queue|bool|
|content_based_deduplication|bool|
|deduplication_scope|utf8|
|fifo_throughput_limit|utf8|
|unknown_fields|json|