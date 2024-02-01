# Table: aws_stepfunctions_map_runs

This table shows data for Stepfunctions Map Runs.

https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeMapRun.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

This table depends on [aws_stepfunctions_executions](aws_stepfunctions_executions.md).

The following tables depend on aws_stepfunctions_map_runs:
  - [aws_stepfunctions_map_run_executions](aws_stepfunctions_map_run_executions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
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
|redrive_count|`int64`|
|redrive_date|`timestamp[us, tz=UTC]`|
|stop_date|`timestamp[us, tz=UTC]`|