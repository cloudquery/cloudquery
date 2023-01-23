# Table: fastly_service_domains

https://developer.fastly.com/reference/api/services/domain/

The composite primary key for this table is (**name**, **service_id**, **service_version**).

## Relations

This table depends on [fastly_service_versions](fastly_service_versions.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|name (PK)|String|
|service_id (PK)|String|
|service_version (PK)|Int|
|comment|String|
|created_at|Timestamp|
|deleted_at|Timestamp|
|updated_at|Timestamp|