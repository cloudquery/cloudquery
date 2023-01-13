# Table: fastly_service_versions

https://developer.fastly.com/reference/api/services/version/

The composite primary key for this table is (**number**, **service_id**).

## Relations

This table depends on [fastly_services](fastly_services.md).

The following tables depend on fastly_service_versions:
  - [fastly_service_backends](fastly_service_backends.md)
  - [fastly_service_domains](fastly_service_domains.md)
  - [fastly_service_health_checks](fastly_service_health_checks.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|number (PK)|Int|
|service_id (PK)|String|
|active|Bool|
|comment|String|
|created_at|Timestamp|
|deleted_at|Timestamp|
|deployed|Bool|
|locked|Bool|
|staging|Bool|
|testing|Bool|
|updated_at|Timestamp|