# Table: aws_glue_databases

This table shows data for Glue Databases.

https://docs.aws.amazon.com/glue/latest/webapi/API_Database.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_glue_databases:
  - [aws_glue_database_tables](aws_glue_database_tables.md)

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
|catalog_id|`utf8`|
|create_table_default_permissions|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|federated_database|`json`|
|location_uri|`utf8`|
|parameters|`json`|
|target_database|`json`|