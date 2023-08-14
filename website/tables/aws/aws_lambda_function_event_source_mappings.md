# Table: aws_lambda_function_event_source_mappings

This table shows data for AWS Lambda Function Event Source Mappings.

https://docs.aws.amazon.com/lambda/latest/dg/API_EventSourceMappingConfiguration.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_lambda_functions](aws_lambda_functions).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|function_arn|`utf8`|
|amazon_managed_kafka_event_source_config|`json`|
|batch_size|`int64`|
|bisect_batch_on_function_error|`bool`|
|destination_config|`json`|
|document_db_event_source_config|`json`|
|event_source_arn|`utf8`|
|filter_criteria|`json`|
|function_response_types|`list<item: utf8, nullable>`|
|last_modified|`timestamp[us, tz=UTC]`|
|last_processing_result|`utf8`|
|maximum_batching_window_in_seconds|`int64`|
|maximum_record_age_in_seconds|`int64`|
|maximum_retry_attempts|`int64`|
|parallelization_factor|`int64`|
|queues|`list<item: utf8, nullable>`|
|scaling_config|`json`|
|self_managed_event_source|`json`|
|self_managed_kafka_event_source_config|`json`|
|source_access_configurations|`json`|
|starting_position|`utf8`|
|starting_position_timestamp|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|state_transition_reason|`utf8`|
|topics|`list<item: utf8, nullable>`|
|tumbling_window_in_seconds|`int64`|
|uuid|`utf8`|