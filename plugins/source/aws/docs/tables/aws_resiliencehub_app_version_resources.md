# Table: aws_resiliencehub_app_version_resources

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_PhysicalResource.html

The primary key for this table is **physical_resource_id**.

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
|logical_resource_id|JSON|
|physical_resource_id (PK)|JSON|
|resource_type|String|
|app_components|JSON|
|resource_name|String|