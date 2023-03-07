# Table: gcp_run_services

https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services#Service

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_run_locations](gcp_run_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|description|String|
|uid|String|
|generation|Int|
|labels|JSON|
|annotations|JSON|
|create_time|Timestamp|
|update_time|Timestamp|
|delete_time|Timestamp|
|expire_time|Timestamp|
|creator|String|
|last_modifier|String|
|client|String|
|client_version|String|
|ingress|String|
|launch_stage|String|
|binary_authorization|JSON|
|template|JSON|
|traffic|JSON|
|observed_generation|Int|
|terminal_condition|JSON|
|conditions|JSON|
|latest_ready_revision|String|
|latest_created_revision|String|
|traffic_statuses|JSON|
|uri|String|
|reconciling|Bool|
|etag|String|