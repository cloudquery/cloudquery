# Table: aws_glue_registry_schema_versions

This table shows data for Glue Registry Schema Versions.

https://docs.aws.amazon.com/glue/latest/webapi/API_GetSchemaVersion.html

The composite primary key for this table is (**registry_schema_arn**, **schema_version_id**).

## Relations

This table depends on [aws_glue_registry_schemas](aws_glue_registry_schemas.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|registry_schema_arn (PK)|`utf8`|
|metadata|`json`|
|created_time|`utf8`|
|data_format|`utf8`|
|schema_arn|`utf8`|
|schema_definition|`utf8`|
|schema_version_id (PK)|`utf8`|
|status|`utf8`|
|version_number|`int64`|