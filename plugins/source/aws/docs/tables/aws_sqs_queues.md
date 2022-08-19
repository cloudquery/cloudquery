
# Table: aws_sqs_queues
Amazon Simple Queue Service
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource|
|region|text|The AWS Region of the resource|
|tags|jsonb||
|url|text|The URL of the Amazon SQS queue|
|approximate_number_of_messages|bigint|The approximate number of messages available for retrieval from the queue|
|approximate_number_of_messages_delayed|bigint|The approximate number of messages in the queue that are delayed and not available for reading immediately|
|approximate_number_of_messages_not_visible|bigint|The approximate number of messages that are in flight|
|created_timestamp|bigint|The time when the queue was created in seconds (epoch time)|
|delay_seconds|bigint|The default delay on the queue in seconds|
|last_modified_timestamp|bigint|The time when the queue was last changed in seconds (epoch time)|
|maximum_message_size|bigint|The limit of how many bytes a message can contain before Amazon SQS rejects it|
|message_retention_period|bigint|The length of time, in seconds, for which Amazon SQS retains a message|
|policy|jsonb|The policy of the queue|
|arn|text|The Amazon resource name (ARN) of the queue|
|receive_message_wait_time_seconds|bigint|The length of time, in seconds, for which the ReceiveMessage action waits for a message to arrive|
|redrive_policy|jsonb|The parameters for the dead-letter queue functionality of the source queue as a JSON object|
|visibility_timeout|bigint|The visibility timeout for the queue|
|kms_master_key_id|text|The ID of an Amazon Web Services managed customer master key (CMK) for Amazon SQS or a custom CMK|
|kms_data_key_reuse_period_seconds|bigint|The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling KMS again|
|sqs_managed_sse_enabled|boolean|True if the queue is using SSE-SQS encryption using SQS owned encryption keys|
|fifo_queue|boolean|True if the queue is FIFO queue|
|content_based_deduplication|boolean|True if content-based deduplication is enabled for the queue|
|deduplication_scope|text|Specifies whether message deduplication occurs at the message group or queue level|
|fifo_throughput_limit|text|Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group|
|redrive_allow_policy|jsonb|The parameters for the dead-letter queue functionality of the source queue as a JSON object|
|unknown_fields|jsonb||
