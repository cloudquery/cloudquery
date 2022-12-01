# Table: aws_athena_data_catalogs

https://docs.aws.amazon.com/athena/latest/APIReference/API_DataCatalog.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_athena_data_catalogs:
  - [aws_athena_data_catalog_databases](aws_athena_data_catalog_databases.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|name|String|
|type|String|
|description|String|
|parameters|JSON|