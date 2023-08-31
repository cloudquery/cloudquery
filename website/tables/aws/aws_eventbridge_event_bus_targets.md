# Table: aws_eventbridge_event_bus_targets

This table shows data for Amazon EventBridge Event Bus Targets.

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Target.html

The composite primary key for this table is (**rule_arn**, **event_bus_arn**, **id**).

## Relations

This table depends on [aws_eventbridge_event_bus_rules](aws_eventbridge_event_bus_rules).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|rule_arn (PK)|`utf8`|
|event_bus_arn (PK)|`utf8`|
|arn|`utf8`|
|id (PK)|`utf8`|
|batch_parameters|`json`|
|dead_letter_config|`json`|
|ecs_parameters|`json`|
|http_parameters|`json`|
|input|`utf8`|
|input_path|`utf8`|
|input_transformer|`json`|
|kinesis_parameters|`json`|
|redshift_data_parameters|`json`|
|retry_policy|`json`|
|role_arn|`utf8`|
|run_command_parameters|`json`|
|sage_maker_pipeline_parameters|`json`|
|sqs_parameters|`json`|