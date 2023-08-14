# Table: aws_glue_database_table_indexes

This table shows data for Glue Database Table Indexes.

https://docs.aws.amazon.com/glue/latest/webapi/API_PartitionIndexDescriptor.html

The composite primary key for this table is (**database_arn**, **database_table_name**, **index_name**).

## Relations

This table depends on [aws_glue_database_tables](aws_glue_database_tables).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|database_arn (PK)|`utf8`|
|database_table_name (PK)|`utf8`|
|index_name (PK)|`utf8`|
|index_status|`utf8`|
|keys|`json`|
|backfill_errors|`json`|