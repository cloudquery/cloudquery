# Table: k8s_core_pvs

This table shows data for Kubernetes (K8s) Core Persistent Volumes (PVs).

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|context|`utf8`|
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
|spec_capacity|`json`|
|spec_persistent_volume_source|`json`|
|spec_access_modes|`list<item: utf8, nullable>`|
|spec_claim_ref|`json`|
|spec_persistent_volume_reclaim_policy|`utf8`|
|spec_storage_class_name|`utf8`|
|spec_mount_options|`list<item: utf8, nullable>`|
|spec_volume_mode|`utf8`|
|spec_node_affinity|`json`|
|status_phase|`utf8`|
|status_message|`utf8`|
|status_reason|`utf8`|