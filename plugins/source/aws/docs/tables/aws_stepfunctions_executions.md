# Table: aws_stepfunctions_executions

This table shows data for Stepfunctions Executions.

https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeExecution.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

This table depends on [aws_stepfunctions_state_machines](aws_stepfunctions_state_machines.md).

The following tables depend on aws_stepfunctions_executions:
  - [aws_stepfunctions_map_runs](aws_stepfunctions_map_runs.md)

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
|start_date|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|cause|`utf8`|
|error|`utf8`|
|input|`utf8`|
|input_details|`json`|
|map_run_arn|`utf8`|
|name|`utf8`|
|output|`utf8`|
|output_details|`json`|
|redrive_count|`int64`|
|redrive_date|`timestamp[us, tz=UTC]`|
|redrive_status|`utf8`|
|redrive_status_reason|`utf8`|
|state_machine_alias_arn|`utf8`|
|state_machine_version_arn|`utf8`|
|stop_date|`timestamp[us, tz=UTC]`|
|trace_header|`utf8`|