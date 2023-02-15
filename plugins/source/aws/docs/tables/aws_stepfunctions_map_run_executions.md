# Table: aws_stepfunctions_map_run_executions

https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeExecution.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_stepfunctions_map_runs](aws_stepfunctions_map_runs.md).

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
|map_run_arn|String|
|state_machine_arn|String|
|execution_arn|String|
|start_date|Timestamp|
|status|String|
|cause|String|
|error|String|
|input|String|
|input_details|JSON|
|name|String|
|output|String|
|output_details|JSON|
|stop_date|Timestamp|
|trace_header|String|