# Table: fastly_service_versions

This table shows data for Fastly Service Versions.

https://developer.fastly.com/reference/api/services/version/

The composite primary key for this table is (**number**, **service_id**).

## Relations

This table depends on [fastly_services](fastly_services).

The following tables depend on fastly_service_versions:
  - [fastly_service_backends](fastly_service_backends)
  - [fastly_service_domains](fastly_service_domains)
  - [fastly_service_health_checks](fastly_service_health_checks)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|number (PK)|`int64`|
|service_id (PK)|`utf8`|
|active|`bool`|
|comment|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|deleted_at|`timestamp[us, tz=UTC]`|
|deployed|`bool`|
|locked|`bool`|
|staging|`bool`|
|testing|`bool`|
|updated_at|`timestamp[us, tz=UTC]`|