# Table: k8s_storage_csi_drivers



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
|spec_attach_required|Bool|
|spec_pod_info_on_mount|Bool|
|spec_volume_lifecycle_modes|StringArray|
|spec_storage_capacity|Bool|
|spec_fs_group_policy|String|
|spec_token_requests|JSON|
|spec_requires_republish|Bool|
|spec_se_linux_mount|Bool|