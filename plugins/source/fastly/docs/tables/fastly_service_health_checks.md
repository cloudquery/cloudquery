# Table: fastly_service_health_checks

https://developer.fastly.com/reference/api/services/healthcheck/

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
|check_interval|Int|
|comment|String|
|created_at|Timestamp|
|deleted_at|Timestamp|
|expected_response|Int|
|http_version|String|
|headers|StringArray|
|host|String|
|initial|Int|
|method|String|
|name (PK)|String|
|path|String|
|service_id (PK)|String|
|service_version (PK)|Int|
|threshold|Int|
|timeout|Int|
|updated_at|Timestamp|
|window|Int|