# Table: fastly_service_health_checks

This table shows data for Fastly Service Health Checks.

https://developer.fastly.com/reference/api/services/healthcheck/

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
|check_interval|`int64`|
|comment|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|deleted_at|`timestamp[us, tz=UTC]`|
|expected_response|`int64`|
|http_version|`utf8`|
|headers|`list<item: utf8, nullable>`|
|host|`utf8`|
|initial|`int64`|
|method|`utf8`|
|path|`utf8`|
|threshold|`int64`|
|timeout|`int64`|
|updated_at|`timestamp[us, tz=UTC]`|
|window|`int64`|