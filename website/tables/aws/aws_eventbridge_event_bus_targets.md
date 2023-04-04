# Table: aws_eventbridge_event_bus_targets

This table shows data for Amazon EventBridge Event Bus Targets.

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Target.html

The composite primary key for this table is (**rule_arn**, **event_bus_arn**, **id**).

## Relations

This table depends on [aws_eventbridge_event_bus_rules](aws_eventbridge_event_bus_rules).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|rule_arn (PK)|String|
|event_bus_arn (PK)|String|
|arn|String|
|id (PK)|String|
|batch_parameters|JSON|
|dead_letter_config|JSON|
|ecs_parameters|JSON|
|http_parameters|JSON|
|input|String|
|input_path|String|
|input_transformer|JSON|
|kinesis_parameters|JSON|
|redshift_data_parameters|JSON|
|retry_policy|JSON|
|role_arn|String|
|run_command_parameters|JSON|
|sage_maker_pipeline_parameters|JSON|
|sqs_parameters|JSON|