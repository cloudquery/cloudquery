# Table: aws_glue_registries

This table shows data for Glue Registries.

https://docs.aws.amazon.com/glue/latest/webapi/API_RegistryListItem.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_glue_registries:
  - [aws_glue_registry_schemas](aws_glue_registry_schemas)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|created_time|`utf8`|
|description|`utf8`|
|registry_arn|`utf8`|
|registry_name|`utf8`|
|status|`utf8`|
|updated_time|`utf8`|