
# Table: aws_sqs_queues
Simple Queue Service
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|url|text|Queue URL|
|policy|jsonb|The queue's policy. A valid Amazon Web Services policy.|
|visibility_timeout|integer|The visibility timeout for the queue, in seconds.|
|maximum_message_size|integer|The limit of how many bytes a message can contain before Amazon SQS rejects it.|
|message_retention_period|integer|The length of time, in seconds, for which Amazon SQS retains a message.|
|approximate_number_of_messages|integer|The approximate number of messages available for retrieval from the queue.|
|approximate_number_of_messages_not_visible|integer|The approximate number of messages that are in flight.|
|created_timestamp|timestamp without time zone|UNIX time when the queue was created.|
|last_modified_timestamp|timestamp without time zone|UNIX time when the queue was last changed.|
|arn|text|Amazon resource name (ARN) of the queue.|
|approximate_number_of_messages_delayed|integer|The approximate number of messages in the queue that are delayed and not available for reading immediately.|
|delay_seconds|integer|The default delay on the queue in seconds.|
|receive_message_wait_time_seconds|integer|the length of time, in seconds, for which the ReceiveMessage action waits for a message to arrive.|
|redrive_policy|jsonb|The parameters for the dead-letter queue functionality of the source queue as a JSON object.|
|fifo_queue|boolean|True if the queue is FIFO queue.|
|content_based_deduplication|boolean|True if content-based deduplication is enabled for the queue.|
|kms_master_key_id|text|ID of an Amazon Web Services managed customer master key (CMK) for Amazon SQS or a custom CMK.|
|kms_data_key_reuse_period_seconds|integer|The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling KMS again.|
|deduplication_scope|text|Specifies whether message deduplication occurs at the message group or queue level.|
|fifo_throughput_limit|text|Specifies whether message deduplication occurs at the message group or queue level.|
|redrive_allow_policy|jsonb|The parameters for the permissions for the dead-letter queue redrive permission.|
|tags|jsonb|Queue tags.|
|unknown_fields|jsonb|Other queue attributes|
