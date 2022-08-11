
# Table: aws_lambda_function_event_source_mappings
A mapping between an Amazon Web Services resource and a Lambda function
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_cq_id|uuid|Unique CloudQuery ID of aws_lambda_functions table (FK)|
|batch_size|integer|The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function|
|bisect_batch_on_function_error|boolean|(Streams only) If the function returns an error, split the batch in two and retry|
|on_failure_destination|text|The Amazon Resource Name (ARN) of the destination resource.|
|on_success_destination|text|The Amazon Resource Name (ARN) of the destination resource.|
|event_source_arn|text|The Amazon Resource Name (ARN) of the event source.|
|criteria_filters|text[]|A list of filters.|
|function_arn|text|The ARN of the Lambda function.|
|function_response_types|text[]|(Streams only) A list of current response type enums applied to the event source mapping.|
|last_modified|timestamp without time zone|The date that the event source mapping was last updated or that its state changed.|
|last_processing_result|text|The result of the last Lambda invocation of your function.|
|maximum_batching_window_in_seconds|integer|(Streams and Amazon SQS standard queues) The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function|
|maximum_record_age_in_seconds|integer|(Streams only) Discard records older than the specified age|
|maximum_retry_attempts|integer|(Streams only) Discard records after the specified number of retries|
|parallelization_factor|integer|(Streams only) The number of batches to process concurrently from each shard. The default value is 1.|
|queues|text[]|(Amazon MQ) The name of the Amazon MQ broker destination queue to consume.|
|self_managed_event_source_endpoints|jsonb|The list of bootstrap servers for your Kafka brokers in the following format: "KAFKA_BOOTSTRAP_SERVERS": ["abc.xyz.com:xxxx","abc2.xyz.com:xxxx"].|
|source_access_configurations|jsonb|An array of the authentication protocol, VPC components, or virtual host to secure and define your event source.|
|starting_position|text|The position in a stream from which to start reading|
|starting_position_timestamp|timestamp without time zone|With StartingPosition set to AT_TIMESTAMP, the time from which to start reading.|
|state|text|The state of the event source mapping|
|state_transition_reason|text|Indicates whether a user or Lambda made the last change to the event source mapping.|
|topics|text[]|The name of the Kafka topic.|
|tumbling_window_in_seconds|integer|(Streams only) The duration in seconds of a processing window|
|uuid|text|The identifier of the event source mapping.|
