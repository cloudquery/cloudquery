# Table: aws_athena_data_catalogs

This table shows data for Athena Data Catalogs.

https://docs.aws.amazon.com/athena/latest/APIReference/API_DataCatalog.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_athena_data_catalogs:
  - [aws_athena_data_catalog_databases](aws_athena_data_catalog_databases)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|name|`utf8`|
|type|`utf8`|
|description|`utf8`|
|parameters|`json`|