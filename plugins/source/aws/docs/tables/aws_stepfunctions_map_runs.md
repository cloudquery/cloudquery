# Table: aws_stepfunctions_map_runs

https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeMapRun.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_stepfunctions_executions](aws_stepfunctions_executions.md).

The following tables depend on aws_stepfunctions_map_runs:
  - [aws_stepfunctions_map_run_executions](aws_stepfunctions_map_run_executions.md)

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
|state_machine_arn|String|
|execution_arn|String|
|execution_counts|JSON|
|item_counts|JSON|
|map_run_arn|String|
|max_concurrency|Int|
|start_date|Timestamp|
|status|String|
|tolerated_failure_count|Int|
|tolerated_failure_percentage|Float|
|stop_date|Timestamp|