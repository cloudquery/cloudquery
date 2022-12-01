# Table: k8s_core_services



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
|spec_cluster_ip|Inet|
|spec_cluster_ips|InetArray|
|spec_external_ips|InetArray|
|spec_load_balancer_ip|Inet|
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
|spec_ports|JSON|
|spec_selector|JSON|
|spec_type|String|
|spec_session_affinity|String|
|spec_load_balancer_source_ranges|StringArray|
|spec_external_name|String|
|spec_external_traffic_policy|String|
|spec_health_check_node_port|Int|
|spec_publish_not_ready_addresses|Bool|
|spec_session_affinity_config|JSON|
|spec_ip_families|StringArray|
|spec_ip_family_policy|String|
|spec_allocate_load_balancer_node_ports|Bool|
|spec_load_balancer_class|String|
|spec_internal_traffic_policy|String|
|status_load_balancer|JSON|
|status_conditions|JSON|