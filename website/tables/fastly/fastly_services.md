# Table: fastly_services

This table shows data for Fastly Services.

https://developer.fastly.com/reference/api/services/service/

The primary key for this table is **id**.

## Relations

The following tables depend on fastly_services:
  - [fastly_service_versions](fastly_service_versions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|active_version|`int64`|
|comment|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|customer_id|`utf8`|
|deleted_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|type|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|versions|`json`|