# Table: aws_glue_database_tables



The composite primary key for this table is (**database_arn**, **name**).

## Relations
This table depends on [aws_glue_databases](aws_glue_databases.md).

The following tables depend on aws_glue_database_tables:
  - [aws_glue_database_table_indexes](aws_glue_database_table_indexes.md)

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
|name (PK)|String|
|catalog_id|String|
|create_time|Timestamp|
|created_by|String|
|database_name|String|
|description|String|
|is_registered_with_lake_formation|Bool|
|last_access_time|Timestamp|
|last_analyzed_time|Timestamp|
|owner|String|
|parameters|JSON|
|partition_keys|JSON|
|retention|Int|
|storage_descriptor|JSON|
|table_type|String|
|target_table|JSON|
|update_time|Timestamp|
|version_id|String|
|view_expanded_text|String|
|view_original_text|String|