# Table: aws_athena_work_group_query_executions

This table shows data for Athena Work Group Query Executions.

https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryExecution.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_athena_work_groups](aws_athena_work_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|work_group_arn|`utf8`|
|engine_version|`json`|
|execution_parameters|`list<item: utf8, nullable>`|
|query|`utf8`|
|query_execution_context|`json`|
|query_execution_id|`utf8`|
|result_configuration|`json`|
|result_reuse_configuration|`json`|
|statement_type|`utf8`|
|statistics|`json`|
|status|`json`|
|substatement_type|`utf8`|
|work_group|`utf8`|