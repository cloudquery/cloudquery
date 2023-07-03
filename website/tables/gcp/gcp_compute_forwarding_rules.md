# Table: gcp_compute_forwarding_rules

This table shows data for GCP Compute Forwarding Rules.

https://cloud.google.com/compute/docs/reference/rest/v1/forwardingRules#ForwardingRule

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|ip_address|`utf8`|
|ip_protocol|`utf8`|
|all_ports|`bool`|
|allow_global_access|`bool`|
|backend_service|`utf8`|
|base_forwarding_rule|`utf8`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|fingerprint|`utf8`|
|id|`int64`|
|ip_version|`utf8`|
|is_mirroring_collector|`bool`|
|kind|`utf8`|
|label_fingerprint|`utf8`|
|labels|`json`|
|load_balancing_scheme|`utf8`|
|metadata_filters|`json`|
|name|`utf8`|
|network|`utf8`|
|network_tier|`utf8`|
|no_automate_dns_zone|`bool`|
|port_range|`utf8`|
|ports|`list<item: utf8, nullable>`|
|psc_connection_id|`int64`|
|psc_connection_status|`utf8`|
|region|`utf8`|
|self_link (PK)|`utf8`|
|service_directory_registrations|`json`|
|service_label|`utf8`|
|service_name|`utf8`|
|source_ip_ranges|`list<item: utf8, nullable>`|
|subnetwork|`utf8`|
|target|`utf8`|