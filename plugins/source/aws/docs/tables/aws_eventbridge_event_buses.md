# Table: aws_eventbridge_event_buses


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_eventbridge_event_buses`:
  - [`aws_eventbridge_event_bus_rules`](aws_eventbridge_event_bus_rules.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|tags|JSON|
|arn (PK)|String|
|name|String|
|policy|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|