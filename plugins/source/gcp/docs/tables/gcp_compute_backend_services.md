# Table: gcp_compute_backend_services



The primary key for this table is **self_link**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|self_link (PK)|String|
|affinity_cookie_ttl_sec|Int|
|backends|JSON|
|cdn_policy|JSON|
|circuit_breakers|JSON|
|compression_mode|String|
|connection_draining|JSON|
|connection_tracking_policy|JSON|
|consistent_hash|JSON|
|creation_timestamp|String|
|custom_request_headers|StringArray|
|custom_response_headers|StringArray|
|description|String|
|edge_security_policy|String|
|enable_cdn|Bool|
|failover_policy|JSON|
|fingerprint|String|
|health_checks|StringArray|
|iap|JSON|
|id|Int|
|kind|String|
|load_balancing_scheme|String|
|locality_lb_policies|JSON|
|locality_lb_policy|String|
|log_config|JSON|
|max_stream_duration|JSON|
|name|String|
|network|String|
|outlier_detection|JSON|
|port|Int|
|port_name|String|
|protocol|String|
|region|String|
|security_policy|String|
|security_settings|JSON|
|service_bindings|StringArray|
|session_affinity|String|
|subsetting|JSON|
|timeout_sec|Int|