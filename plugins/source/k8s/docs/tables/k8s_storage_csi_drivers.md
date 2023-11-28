# Table: k8s_storage_csi_drivers

This table shows data for Kubernetes (K8s) Storage Container Storage Interface (CSI) Drivers.

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
|spec_attach_required|`bool`|
|spec_pod_info_on_mount|`bool`|
|spec_volume_lifecycle_modes|`list<item: utf8, nullable>`|
|spec_storage_capacity|`bool`|
|spec_fs_group_policy|`utf8`|
|spec_token_requests|`json`|
|spec_requires_republish|`bool`|
|spec_se_linux_mount|`bool`|