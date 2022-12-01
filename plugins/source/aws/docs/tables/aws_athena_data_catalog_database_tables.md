# Table: aws_athena_data_catalog_database_tables

https://docs.aws.amazon.com/athena/latest/APIReference/API_TableMetadata.html

The composite primary key for this table is (**data_catalog_arn**, **data_catalog_database_name**, **name**).

## Relations
This table depends on [aws_athena_data_catalog_databases](aws_athena_data_catalog_databases.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|data_catalog_arn (PK)|String|
|data_catalog_database_name (PK)|String|
|name (PK)|String|
|columns|JSON|
|create_time|Timestamp|
|last_access_time|Timestamp|
|parameters|JSON|
|partition_keys|JSON|
|table_type|String|