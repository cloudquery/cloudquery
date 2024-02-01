# Table: aws_glue_registries

This table shows data for Glue Registries.

https://docs.aws.amazon.com/glue/latest/webapi/API_RegistryListItem.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_glue_registries:
  - [aws_glue_registry_schemas](aws_glue_registry_schemas.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|created_time|`utf8`|
|description|`utf8`|
|registry_arn|`utf8`|
|registry_name|`utf8`|
|status|`utf8`|
|updated_time|`utf8`|