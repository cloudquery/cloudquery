# Table: fastly_services

https://developer.fastly.com/reference/api/services/service/

The primary key for this table is **id**.

## Relations

The following tables depend on fastly_services:
  - [fastly_service_versions](fastly_service_versions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|active_version|Int|
|comment|String|
|created_at|Timestamp|
|customer_id|String|
|deleted_at|Timestamp|
|id (PK)|String|
|name|String|
|type|String|
|updated_at|Timestamp|
|versions|JSON|