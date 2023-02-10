# Table: aws_stepfunctions_executions

https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeStateMachine.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_stepfunctions_executions:
  - [aws_stepfunctions_executions_map_runs](aws_stepfunctions_executions_map_runs.md)

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
|name|String|
|start_date|Timestamp|
|state_machine_arn|String|
|status|String|
|item_count|Int|
|map_run_arn|String|
|stop_date|Timestamp|