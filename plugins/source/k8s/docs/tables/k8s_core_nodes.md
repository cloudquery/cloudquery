# Table: k8s_core_nodes

This table shows data for Kubernetes (K8s) Core Nodes.

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|context|`utf8`|
|spec_pod_cidr|`inet`|
|spec_pod_cidrs|`list<item: inet, nullable>`|
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
|spec_provider_id|`utf8`|
|spec_unschedulable|`bool`|
|spec_taints|`json`|
|spec_config_source|`json`|
|status_capacity|`json`|
|status_allocatable|`json`|
|status_phase|`utf8`|
|status_conditions|`json`|
|status_addresses|`json`|
|status_daemon_endpoints|`json`|
|status_node_info|`json`|
|status_images|`json`|
|status_volumes_in_use|`list<item: utf8, nullable>`|
|status_volumes_attached|`json`|
|status_config|`json`|