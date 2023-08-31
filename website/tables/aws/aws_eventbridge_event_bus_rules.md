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
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|event_bus_arn|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|description|`utf8`|
|event_bus_name|`utf8`|
|event_pattern|`utf8`|
|managed_by|`utf8`|
|name|`utf8`|
|role_arn|`utf8`|
|schedule_expression|`utf8`|
|state|`utf8`|