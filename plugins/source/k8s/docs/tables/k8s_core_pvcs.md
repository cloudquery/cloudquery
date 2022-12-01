# Table: k8s_core_pvcs



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
|spec_access_modes|StringArray|
|spec_selector|JSON|
|spec_resources|JSON|
|spec_volume_name|String|
|spec_storage_class_name|String|
|spec_volume_mode|String|
|spec_data_source|JSON|
|spec_data_source_ref|JSON|
|status_phase|String|
|status_access_modes|StringArray|
|status_capacity|JSON|
|status_conditions|JSON|
|status_allocated_resources|JSON|
|status_resize_status|String|