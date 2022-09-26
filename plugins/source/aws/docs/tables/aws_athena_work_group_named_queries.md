# Table: aws_athena_work_group_named_queries


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_athena_work_groups`](aws_athena_work_groups.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|work_group_arn|String|
|database|String|
|name|String|
|query_string|String|
|description|String|
|named_query_id|String|
|work_group|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|