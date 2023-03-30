# Table: aws_eventbridge_event_bus_rules

This table shows data for Amazon EventBridge Event Bus Rules.

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Rule.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_eventbridge_event_buses](aws_eventbridge_event_buses).

The following tables depend on aws_eventbridge_event_bus_rules:
  - [aws_eventbridge_event_bus_targets](aws_eventbridge_event_bus_targets)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|event_bus_arn|String|
|tags|JSON|
|arn (PK)|String|
|description|String|
|event_bus_name|String|
|event_pattern|String|
|managed_by|String|
|name|String|
|role_arn|String|
|schedule_expression|String|
|state|String|