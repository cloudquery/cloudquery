# Table: aws_athena_work_group_query_executions

https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryExecution.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_athena_work_groups](aws_athena_work_groups.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|work_group_arn|String|
|engine_version|JSON|
|execution_parameters|StringArray|
|query|String|
|query_execution_context|JSON|
|query_execution_id|String|
|result_configuration|JSON|
|result_reuse_configuration|JSON|
|statement_type|String|
|statistics|JSON|
|status|JSON|
|work_group|String|