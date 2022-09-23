# Table: aws_glue_registries


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_glue_registries`:
  - [`aws_glue_registry_schemas`](aws_glue_registry_schemas.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|tags|JSON|
|arn (PK)|String|
|created_time|String|
|description|String|
|registry_name|String|
|status|String|
|updated_time|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|