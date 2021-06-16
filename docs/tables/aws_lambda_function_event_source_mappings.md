
# Table: aws_lambda_function_event_source_mappings
A mapping between an AWS resource and an AWS Lambda function
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_id|uuid|Unique ID of aws_lambda_functions table (FK)|
|batch_size|integer|The maximum number of items to retrieve in a single batch.|
|bisect_batch_on_function_error|boolean|(Streams) If the function returns an error, split the batch in two and retry. The default value is false.|
|on_failure_destination|text|The Amazon Resource Name (ARN) of the destination resource.|
|on_success_destination|text|The Amazon Resource Name (ARN) of the destination resource.|
|event_source_arn|text|The Amazon Resource Name (ARN) of the event source.|
|function_arn|text|The ARN of the Lambda function.|
|function_response_types|text[]|(Streams) A list of current response type enums applied to the event source mapping.|
|last_modified|timestamp without time zone|The date that the event source mapping was last updated, or its state changed.|
|last_processing_result|text|The result of the last AWS Lambda invocation of your Lambda function.|
|maximum_batching_window_in_seconds|integer|(Streams and SQS standard queues) The maximum amount of time to gather records before invoking the function, in seconds|
|maximum_record_age_in_seconds|integer|(Streams) Discard records older than the specified age|
|maximum_retry_attempts|integer|(Streams) Discard records after the specified number of retries|
|parallelization_factor|integer|(Streams) The number of batches to process from each shard concurrently|
|queues|text[]|(MQ) The name of the Amazon MQ broker destination queue to consume.|
|self_managed_event_source_endpoints|jsonb|The list of bootstrap servers for your Kafka brokers in the following format: "KAFKA_BOOTSTRAP_SERVERS": ["abc.xyz.com:xxxx","abc2.xyz.com:xxxx"].|
|starting_position|text|The position in a stream from which to start reading|
|starting_position_timestamp|timestamp without time zone|With StartingPosition set to AT_TIMESTAMP, the time from which to start reading.|
|state|text|The state of the event source mapping|
|state_transition_reason|text|Indicates whether the last change to the event source mapping was made by a user, or by the Lambda service.|
|topics|text[]|The name of the Kafka topic.|
|tumbling_window_in_seconds|integer|(Streams) The duration in seconds of a processing window|
|uuid|text|The identifier of the event source mapping.|
