# Table: gcp_compute_projects

This table shows data for GCP Compute Projects.

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|project_id|utf8|
|common_instance_metadata|json|
|creation_timestamp|utf8|
|default_network_tier|utf8|
|default_service_account|utf8|
|description|utf8|
|enabled_features|list<item: utf8, nullable>|
|id|int64|
|kind|utf8|
|name|utf8|
|quotas|json|
|self_link (PK)|utf8|
|usage_export_location|json|
|vm_dns_setting|utf8|
|xpn_project_status|utf8|