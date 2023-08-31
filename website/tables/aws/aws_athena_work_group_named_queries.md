# Table: aws_athena_work_group_named_queries

This table shows data for Athena Work Group Named Queries.

https://docs.aws.amazon.com/athena/latest/APIReference/API_NamedQuery.html

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
|database|`utf8`|
|name|`utf8`|
|query_string|`utf8`|
|description|`utf8`|
|named_query_id|`utf8`|
|work_group|`utf8`|