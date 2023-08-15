# Table: aws_glue_registry_schemas

This table shows data for Glue Registry Schemas.

https://docs.aws.amazon.com/glue/latest/webapi/API_GetSchema.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_glue_registries](aws_glue_registries).

The following tables depend on aws_glue_registry_schemas:
  - [aws_glue_registry_schema_versions](aws_glue_registry_schema_versions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|compatibility|`utf8`|
|created_time|`utf8`|
|data_format|`utf8`|
|description|`utf8`|
|latest_schema_version|`int64`|
|next_schema_version|`int64`|
|registry_arn|`utf8`|
|registry_name|`utf8`|
|schema_arn|`utf8`|
|schema_checkpoint|`int64`|
|schema_name|`utf8`|
|schema_status|`utf8`|
|updated_time|`utf8`|
|result_metadata|`json`|