# Table: aws_athena_work_group_prepared_statements


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_athena_work_groups`](aws_athena_work_groups.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|work_group_arn|String|
|description|String|
|last_modified_time|Timestamp|
|query_statement|String|
|statement_name|String|
|work_group_name|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|