# Table: aws_stepfunctions_executions_map_runs

https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeStateMachine.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_stepfunctions_executions](aws_stepfunctions_executions.md).

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
|execution_arn|String|
|map_run_arn|String|
|start_date|Timestamp|
|state_machine_arn|String|
|stop_date|Timestamp|