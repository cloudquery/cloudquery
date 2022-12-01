# Table: aws_athena_work_groups

https://docs.aws.amazon.com/athena/latest/APIReference/API_WorkGroup.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_athena_work_groups:
  - [aws_athena_work_group_prepared_statements](aws_athena_work_group_prepared_statements.md)
  - [aws_athena_work_group_query_executions](aws_athena_work_group_query_executions.md)
  - [aws_athena_work_group_named_queries](aws_athena_work_group_named_queries.md)

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
|tags|JSON|
|name|String|
|configuration|JSON|
|creation_time|Timestamp|
|description|String|
|state|String|