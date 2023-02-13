# Table: aws_resiliencehub_app_version_resource_mappings

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_ResourceMapping.html

The composite primary key for this table is (**app_arn**, **app_version**, **physical_resource_identifier**).

## Relations

This table depends on [aws_resiliencehub_app_versions](aws_resiliencehub_app_versions.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|app_arn (PK)|String|
|app_version (PK)|String|
|physical_resource_identifier (PK)|String|
|mapping_type|String|
|physical_resource_id|JSON|
|app_registry_app_name|String|
|logical_stack_name|String|
|resource_group_name|String|
|resource_name|String|
|terraform_source_name|String|