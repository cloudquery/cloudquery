# Table: aws_stepfunctions_state_machines

This table shows data for Stepfunctions State Machines.

https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeStateMachine.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_stepfunctions_state_machines:
  - [aws_stepfunctions_executions](aws_stepfunctions_executions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|creation_date|`timestamp[us, tz=UTC]`|
|definition|`utf8`|
|name|`utf8`|
|role_arn|`utf8`|
|state_machine_arn|`utf8`|
|type|`utf8`|
|description|`utf8`|
|label|`utf8`|
|logging_configuration|`json`|
|revision_id|`utf8`|
|status|`utf8`|
|tracing_configuration|`json`|