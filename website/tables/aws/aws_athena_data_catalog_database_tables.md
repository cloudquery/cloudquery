# Table: aws_athena_data_catalog_database_tables

This table shows data for Athena Data Catalog Database Tables.

https://docs.aws.amazon.com/athena/latest/APIReference/API_TableMetadata.html

The composite primary key for this table is (**data_catalog_arn**, **data_catalog_database_name**, **name**).

## Relations

This table depends on [aws_athena_data_catalog_databases](aws_athena_data_catalog_databases).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|data_catalog_arn (PK)|`utf8`|
|data_catalog_database_name (PK)|`utf8`|
|name (PK)|`utf8`|
|columns|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|last_access_time|`timestamp[us, tz=UTC]`|
|parameters|`json`|
|partition_keys|`json`|
|table_type|`utf8`|