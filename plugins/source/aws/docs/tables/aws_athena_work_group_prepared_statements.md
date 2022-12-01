# Table: aws_athena_work_group_prepared_statements

https://docs.aws.amazon.com/athena/latest/APIReference/API_PreparedStatement.html

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
|description|String|
|last_modified_time|Timestamp|
|query_statement|String|
|statement_name|String|
|work_group_name|String|