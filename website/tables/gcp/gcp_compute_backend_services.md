# Table: gcp_compute_backend_services

This table shows data for GCP Compute Backend Services.

https://cloud.google.com/compute/docs/reference/rest/v1/backendServices#BackendService

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|affinity_cookie_ttl_sec|`int64`|
|backends|`json`|
|cdn_policy|`json`|
|circuit_breakers|`json`|
|compression_mode|`utf8`|
|connection_draining|`json`|
|connection_tracking_policy|`json`|
|consistent_hash|`json`|
|creation_timestamp|`utf8`|
|custom_request_headers|`list<item: utf8, nullable>`|
|custom_response_headers|`list<item: utf8, nullable>`|
|description|`utf8`|
|edge_security_policy|`utf8`|
|enable_cdn|`bool`|
|failover_policy|`json`|
|fingerprint|`utf8`|
|health_checks|`list<item: utf8, nullable>`|
|iap|`json`|
|id|`int64`|
|kind|`utf8`|
|load_balancing_scheme|`utf8`|
|locality_lb_policies|`json`|
|locality_lb_policy|`utf8`|
|log_config|`json`|
|max_stream_duration|`json`|
|name|`utf8`|
|network|`utf8`|
|outlier_detection|`json`|
|port|`int64`|
|port_name|`utf8`|
|protocol|`utf8`|
|region|`utf8`|
|security_policy|`utf8`|
|security_settings|`json`|
|self_link (PK)|`utf8`|
|service_bindings|`list<item: utf8, nullable>`|
|session_affinity|`utf8`|
|subsetting|`json`|
|timeout_sec|`int64`|