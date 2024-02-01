# Table: aws_glue_database_table_indexes

This table shows data for Glue Database Table Indexes.

https://docs.aws.amazon.com/glue/latest/webapi/API_PartitionIndexDescriptor.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**database_arn**, **database_table_name**, **index_name**).
## Relations

This table depends on [aws_glue_database_tables](aws_glue_database_tables.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|database_arn|`utf8`|
|database_table_name|`utf8`|
|index_name|`utf8`|
|index_status|`utf8`|
|keys|`json`|
|backfill_errors|`json`|