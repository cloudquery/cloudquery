# Table: aws_eventbridge_event_buses

This table shows data for Amazon EventBridge Event Buses.

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_EventBus.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_eventbridge_event_buses:
  - [aws_eventbridge_event_bus_rules](aws_eventbridge_event_bus_rules)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|name|`utf8`|
|policy|`utf8`|