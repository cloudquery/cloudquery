# Table: gcp_compute_forwarding_rules



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
|ip_address|String|
|ip_protocol|String|
|all_ports|Bool|
|allow_global_access|Bool|
|backend_service|String|
|creation_timestamp|String|
|description|String|
|fingerprint|String|
|id|Int|
|ip_version|String|
|is_mirroring_collector|Bool|
|kind|String|
|label_fingerprint|String|
|labels|JSON|
|load_balancing_scheme|String|
|metadata_filters|JSON|
|name|String|
|network|String|
|network_tier|String|
|no_automate_dns_zone|Bool|
|port_range|String|
|ports|StringArray|
|psc_connection_id|Int|
|psc_connection_status|String|
|region|String|
|service_directory_registrations|JSON|
|service_label|String|
|service_name|String|
|subnetwork|String|
|target|String|