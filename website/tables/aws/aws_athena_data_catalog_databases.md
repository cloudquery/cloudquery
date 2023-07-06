# Table: aws_athena_data_catalog_databases

This table shows data for Athena Data Catalog Databases.

https://docs.aws.amazon.com/athena/latest/APIReference/API_Database.html

The composite primary key for this table is (**data_catalog_arn**, **name**).

## Relations

This table depends on [aws_athena_data_catalogs](aws_athena_data_catalogs).

The following tables depend on aws_athena_data_catalog_databases:
  - [aws_athena_data_catalog_database_tables](aws_athena_data_catalog_database_tables)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|data_catalog_arn (PK)|`utf8`|
|name (PK)|`utf8`|
|description|`utf8`|
|parameters|`json`|