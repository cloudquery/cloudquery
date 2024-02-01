# Table: aws_athena_data_catalog_database_tables

This table shows data for Athena Data Catalog Database Tables.

https://docs.aws.amazon.com/athena/latest/APIReference/API_TableMetadata.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**data_catalog_arn**, **data_catalog_database_name**, **name**).
## Relations

This table depends on [aws_athena_data_catalog_databases](aws_athena_data_catalog_databases.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|data_catalog_arn|`utf8`|
|data_catalog_database_name|`utf8`|
|name|`utf8`|
|columns|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|last_access_time|`timestamp[us, tz=UTC]`|
|parameters|`json`|
|partition_keys|`json`|
|table_type|`utf8`|