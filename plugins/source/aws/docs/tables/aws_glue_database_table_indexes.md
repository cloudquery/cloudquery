# Table: aws_glue_database_table_indexes



The composite primary key for this table is (**database_arn**, **database_table_name**, **index_name**).

## Relations
This table depends on [aws_glue_database_tables](aws_glue_database_tables.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|database_arn (PK)|String|
|database_table_name (PK)|String|
|index_name (PK)|String|
|index_status|String|
|keys|JSON|
|backfill_errors|JSON|