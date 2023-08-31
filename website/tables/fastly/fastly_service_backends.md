# Table: fastly_service_backends

This table shows data for Fastly Service Backends.

https://developer.fastly.com/reference/api/services/backend/

The composite primary key for this table is (**name**, **service_id**, **service_version**).

## Relations

This table depends on [fastly_service_versions](fastly_service_versions).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|
|sslca_cert|`utf8`|
|service_id (PK)|`utf8`|
|service_version (PK)|`int64`|
|address|`utf8`|
|auto_loadbalance|`bool`|
|between_bytes_timeout|`int64`|
|comment|`utf8`|
|connect_timeout|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|deleted_at|`timestamp[us, tz=UTC]`|
|error_threshold|`int64`|
|first_byte_timeout|`int64`|
|health_check|`utf8`|
|hostname|`utf8`|
|max_conn|`int64`|
|max_tls_version|`utf8`|
|min_tls_version|`utf8`|
|override_host|`utf8`|
|port|`int64`|
|request_condition|`utf8`|
|ssl_cert_hostname|`utf8`|
|ssl_check_cert|`bool`|
|ssl_ciphers|`utf8`|
|ssl_client_cert|`utf8`|
|ssl_client_key|`utf8`|
|ssl_hostname|`utf8`|
|sslsni_hostname|`utf8`|
|shield|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|use_ssl|`bool`|
|weight|`int64`|