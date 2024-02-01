# Table: aws_glue_database_tables

This table shows data for Glue Database Tables.

https://docs.aws.amazon.com/glue/latest/webapi/API_Table.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**database_arn**, **name**).
## Relations

This table depends on [aws_glue_databases](aws_glue_databases.md).

The following tables depend on aws_glue_database_tables:
  - [aws_glue_database_table_indexes](aws_glue_database_table_indexes.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|database_arn|`utf8`|
|name|`utf8`|
|catalog_id|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|created_by|`utf8`|
|database_name|`utf8`|
|description|`utf8`|
|federated_table|`json`|
|is_registered_with_lake_formation|`bool`|
|last_access_time|`timestamp[us, tz=UTC]`|
|last_analyzed_time|`timestamp[us, tz=UTC]`|
|owner|`utf8`|
|parameters|`json`|
|partition_keys|`json`|
|retention|`int64`|
|storage_descriptor|`json`|
|table_type|`utf8`|
|target_table|`json`|
|update_time|`timestamp[us, tz=UTC]`|
|version_id|`utf8`|
|view_expanded_text|`utf8`|
|view_original_text|`utf8`|