# Table: aws_stepfunctions_map_runs

This table shows data for Stepfunctions Map Runs.

https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeMapRun.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_stepfunctions_executions](aws_stepfunctions_executions).

The following tables depend on aws_stepfunctions_map_runs:
  - [aws_stepfunctions_map_run_executions](aws_stepfunctions_map_run_executions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|state_machine_arn|`utf8`|
|execution_arn|`utf8`|
|execution_counts|`json`|
|item_counts|`json`|
|map_run_arn|`utf8`|
|max_concurrency|`int64`|
|start_date|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|tolerated_failure_count|`int64`|
|tolerated_failure_percentage|`float64`|
|stop_date|`timestamp[us, tz=UTC]`|