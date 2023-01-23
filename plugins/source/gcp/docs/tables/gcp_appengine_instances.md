# Table: gcp_appengine_instances

https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions.instances#Instance

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_appengine_versions](gcp_appengine_versions.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|id|String|
|app_engine_release|String|
|availability|String|
|vm_name|String|
|vm_zone_name|String|
|vm_id|String|
|start_time|Timestamp|
|requests|Int|
|errors|Int|
|qps|Float|
|average_latency|Int|
|memory_usage|Int|
|vm_status|String|
|vm_debug_enabled|Bool|
|vm_ip|String|
|vm_liveness|String|