# Table: fastly_service_backends

https://developer.fastly.com/reference/api/services/backend/

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
|sslca_cert|String|
|service_id (PK)|String|
|service_version (PK)|Int|
|address|String|
|auto_loadbalance|Bool|
|between_bytes_timeout|Int|
|comment|String|
|connect_timeout|Int|
|created_at|Timestamp|
|deleted_at|Timestamp|
|error_threshold|Int|
|first_byte_timeout|Int|
|health_check|String|
|hostname|String|
|max_conn|Int|
|max_tls_version|String|
|min_tls_version|String|
|override_host|String|
|port|Int|
|request_condition|String|
|ssl_cert_hostname|String|
|ssl_check_cert|Bool|
|ssl_ciphers|String|
|ssl_client_cert|String|
|ssl_client_key|String|
|ssl_hostname|String|
|sslsni_hostname|String|
|shield|String|
|updated_at|Timestamp|
|use_ssl|Bool|
|weight|Int|