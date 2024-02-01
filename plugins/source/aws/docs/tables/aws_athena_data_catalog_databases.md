# Table: aws_athena_data_catalog_databases

This table shows data for Athena Data Catalog Databases.

https://docs.aws.amazon.com/athena/latest/APIReference/API_Database.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**data_catalog_arn**, **name**).
## Relations

This table depends on [aws_athena_data_catalogs](aws_athena_data_catalogs.md).

The following tables depend on aws_athena_data_catalog_databases:
  - [aws_athena_data_catalog_database_tables](aws_athena_data_catalog_database_tables.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|data_catalog_arn|`utf8`|
|name|`utf8`|
|description|`utf8`|
|parameters|`json`|