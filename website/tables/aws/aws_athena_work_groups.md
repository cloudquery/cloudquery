# Table: aws_athena_work_groups

This table shows data for Athena Work Groups.

https://docs.aws.amazon.com/athena/latest/APIReference/API_WorkGroup.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_athena_work_groups:
  - [aws_athena_work_group_named_queries](aws_athena_work_group_named_queries)
  - [aws_athena_work_group_prepared_statements](aws_athena_work_group_prepared_statements)
  - [aws_athena_work_group_query_executions](aws_athena_work_group_query_executions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|name|`utf8`|
|configuration|`json`|
|creation_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|state|`utf8`|