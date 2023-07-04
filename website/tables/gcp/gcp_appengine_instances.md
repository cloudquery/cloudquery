# Table: gcp_appengine_instances

This table shows data for GCP App Engine Instances.

https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions.instances#Instance

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_appengine_versions](gcp_appengine_versions).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|id|`utf8`|
|app_engine_release|`utf8`|
|availability|`utf8`|
|vm_name|`utf8`|
|vm_zone_name|`utf8`|
|vm_id|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|requests|`int64`|
|errors|`int64`|
|qps|`float64`|
|average_latency|`int64`|
|memory_usage|`int64`|
|vm_status|`utf8`|
|vm_debug_enabled|`bool`|
|vm_ip|`utf8`|
|vm_liveness|`utf8`|