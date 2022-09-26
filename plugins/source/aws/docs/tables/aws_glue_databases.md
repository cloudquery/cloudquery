# Table: aws_glue_databases


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_glue_databases`:
  - [`aws_glue_database_tables`](aws_glue_database_tables.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|name|String|
|catalog_id|String|
|create_table_default_permissions|JSON|
|create_time|Timestamp|
|description|String|
|location_uri|String|
|parameters|JSON|
|target_database|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|