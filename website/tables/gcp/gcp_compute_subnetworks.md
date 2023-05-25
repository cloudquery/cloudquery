# Table: gcp_compute_subnetworks

This table shows data for GCP Compute Subnetworks.

https://cloud.google.com/compute/docs/reference/rest/v1/subnetworks#Subnetwork

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|project_id|utf8|
|creation_timestamp|utf8|
|description|utf8|
|enable_flow_logs|bool|
|external_ipv6_prefix|utf8|
|fingerprint|utf8|
|gateway_address|utf8|
|id|int64|
|internal_ipv6_prefix|utf8|
|ip_cidr_range|utf8|
|ipv6_access_type|utf8|
|ipv6_cidr_range|utf8|
|kind|utf8|
|log_config|json|
|name|utf8|
|network|utf8|
|private_ip_google_access|bool|
|private_ipv6_google_access|utf8|
|purpose|utf8|
|region|utf8|
|role|utf8|
|secondary_ip_ranges|json|
|self_link (PK)|utf8|
|stack_type|utf8|
|state|utf8|