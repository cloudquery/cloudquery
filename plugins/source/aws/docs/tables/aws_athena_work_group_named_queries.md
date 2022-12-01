# Table: aws_athena_work_group_named_queries

https://docs.aws.amazon.com/athena/latest/APIReference/API_NamedQuery.html

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
|database|String|
|name|String|
|query_string|String|
|description|String|
|named_query_id|String|
|work_group|String|