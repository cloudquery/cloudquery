# Table: aws_athena_work_group_query_executions


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_athena_work_groups`](aws_athena_work_groups.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|work_group_arn|String|
|engine_version|JSON|
|execution_parameters|StringArray|
|query|String|
|query_execution_context|JSON|
|query_execution_id|String|
|result_configuration|JSON|
|statement_type|String|
|statistics|JSON|
|status|JSON|
|work_group|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|