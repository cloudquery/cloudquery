# Table: k8s_core_nodes



The primary key for this table is **uid**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|context|String|
|uid (PK)|String|
|spec_pod_cidr|CIDR|
|spec_pod_cidrs|CIDRArray|
|kind|String|
|api_version|String|
|name|String|
|namespace|String|
|resource_version|String|
|generation|Int|
|deletion_grace_period_seconds|Int|
|labels|JSON|
|annotations|JSON|
|owner_references|JSON|
|finalizers|StringArray|
|spec_provider_id|String|
|spec_unschedulable|Bool|
|spec_taints|JSON|
|spec_config_source|JSON|
|status_capacity|JSON|
|status_allocatable|JSON|
|status_phase|String|
|status_conditions|JSON|
|status_addresses|JSON|
|status_daemon_endpoints|JSON|
|status_node_info|JSON|
|status_images|JSON|
|status_volumes_in_use|StringArray|
|status_volumes_attached|JSON|
|status_config|JSON|