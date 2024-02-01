# Table: aws_athena_data_catalogs

This table shows data for Athena Data Catalogs.

https://docs.aws.amazon.com/athena/latest/APIReference/API_DataCatalog.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_athena_data_catalogs:
  - [aws_athena_data_catalog_databases](aws_athena_data_catalog_databases.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|name|`utf8`|
|type|`utf8`|
|description|`utf8`|
|parameters|`json`|