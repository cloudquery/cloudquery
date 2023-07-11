# Table: aws_stepfunctions_map_run_executions

This table shows data for Stepfunctions Map Run Executions.

https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeExecution.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_stepfunctions_map_runs](aws_stepfunctions_map_runs).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|map_run_arn|`utf8`|
|state_machine_arn|`utf8`|
|execution_arn|`utf8`|
|start_date|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|cause|`utf8`|
|error|`utf8`|
|input|`utf8`|
|input_details|`json`|
|name|`utf8`|
|output|`utf8`|
|output_details|`json`|
|state_machine_alias_arn|`utf8`|
|state_machine_version_arn|`utf8`|
|stop_date|`timestamp[us, tz=UTC]`|
|trace_header|`utf8`|