# Table: aws_glue_database_table_indexes


The composite primary key for this table is (**database_arn**, **database_table_name**, **index_name**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|database_arn (PK)|String|
|database_table_name (PK)|String|
|index_name (PK)|String|
|index_status|String|
|keys|JSON|
|backfill_errors|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|