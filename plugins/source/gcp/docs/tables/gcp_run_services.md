# Table: gcp_run_services


The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|project_id|String|
|name|String|
|description|String|
|uid|String|
|generation|Int|
|labels|JSON|
|annotations|JSON|
|create_time|JSON|
|update_time|JSON|
|delete_time|JSON|
|expire_time|JSON|
|creator|String|
|last_modifier|String|
|client|String|
|client_version|String|
|ingress|Int|
|launch_stage|Int|
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
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|