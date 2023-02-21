# Table: aws_lambda_function_event_source_mappings

https://docs.aws.amazon.com/lambda/latest/dg/API_EventSourceMappingConfiguration.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_lambda_functions](aws_lambda_functions.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|function_arn|String|
|amazon_managed_kafka_event_source_config|JSON|
|batch_size|Int|
|bisect_batch_on_function_error|Bool|
|destination_config|JSON|
|event_source_arn|String|
|filter_criteria|JSON|
|function_response_types|StringArray|
|last_modified|Timestamp|
|last_processing_result|String|
|maximum_batching_window_in_seconds|Int|
|maximum_record_age_in_seconds|Int|
|maximum_retry_attempts|Int|
|parallelization_factor|Int|
|queues|StringArray|
|scaling_config|JSON|
|self_managed_event_source|JSON|
|self_managed_kafka_event_source_config|JSON|
|source_access_configurations|JSON|
|starting_position|String|
|starting_position_timestamp|Timestamp|
|state|String|
|state_transition_reason|String|
|topics|StringArray|
|tumbling_window_in_seconds|Int|
|uuid|String|