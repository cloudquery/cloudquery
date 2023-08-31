# Table: fastly_service_domains

This table shows data for Fastly Service Domains.

https://developer.fastly.com/reference/api/services/domain/

The composite primary key for this table is (**name**, **service_id**, **service_version**).

## Relations

This table depends on [fastly_service_versions](fastly_service_versions).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|
|service_id (PK)|`utf8`|
|service_version (PK)|`int64`|
|comment|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|deleted_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|