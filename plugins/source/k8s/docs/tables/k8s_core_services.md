# Table: k8s_core_services

This table shows data for Kubernetes (K8s) Core Services.

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|context|`utf8`|
|spec_cluster_ip|`inet`|
|spec_cluster_ips|`list<item: inet, nullable>`|
|spec_external_ips|`list<item: inet, nullable>`|
|spec_load_balancer_ip|`inet`|
|kind|`utf8`|
|api_version|`utf8`|
|name|`utf8`|
|namespace|`utf8`|
|uid (PK)|`utf8`|
|resource_version|`utf8`|
|generation|`int64`|
|deletion_grace_period_seconds|`int64`|
|labels|`json`|
|annotations|`json`|
|owner_references|`json`|
|finalizers|`list<item: utf8, nullable>`|
|spec_ports|`json`|
|spec_selector|`json`|
|spec_type|`utf8`|
|spec_session_affinity|`utf8`|
|spec_load_balancer_source_ranges|`list<item: utf8, nullable>`|
|spec_external_name|`utf8`|
|spec_external_traffic_policy|`utf8`|
|spec_health_check_node_port|`int64`|
|spec_publish_not_ready_addresses|`bool`|
|spec_session_affinity_config|`json`|
|spec_ip_families|`list<item: utf8, nullable>`|
|spec_ip_family_policy|`utf8`|
|spec_allocate_load_balancer_node_ports|`bool`|
|spec_load_balancer_class|`utf8`|
|spec_internal_traffic_policy|`utf8`|
|status_load_balancer|`json`|
|status_conditions|`json`|