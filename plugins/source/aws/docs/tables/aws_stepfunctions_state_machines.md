# Table: aws_stepfunctions_state_machines

https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeStateMachine.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_stepfunctions_state_machines:
  - [aws_stepfunctions_executions](aws_stepfunctions_executions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|creation_date|Timestamp|
|definition|String|
|name|String|
|role_arn|String|
|state_machine_arn|String|
|type|String|
|label|String|
|logging_configuration|JSON|
|status|String|
|tracing_configuration|JSON|