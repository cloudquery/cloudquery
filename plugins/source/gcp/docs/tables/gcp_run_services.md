# Table: gcp_run_services



The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|name|String|
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