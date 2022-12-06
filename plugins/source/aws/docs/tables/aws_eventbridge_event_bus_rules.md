# Table: aws_eventbridge_event_bus_rules

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Rule.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_eventbridge_event_buses](aws_eventbridge_event_buses.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|event_bus_arn|String|
|tags|JSON|
|arn|String|
|description|String|
|event_bus_name|String|
|event_pattern|String|
|managed_by|String|
|name|String|
|role_arn|String|
|schedule_expression|String|
|state|String|