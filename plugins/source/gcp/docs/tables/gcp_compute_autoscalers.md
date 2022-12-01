# Table: gcp_compute_autoscalers



The primary key for this table is **self_link**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|self_link (PK)|String|
|autoscaling_policy|JSON|
|creation_timestamp|String|
|description|String|
|id|Int|
|kind|String|
|name|String|
|recommended_size|Int|
|region|String|
|scaling_schedule_status|JSON|
|status|String|
|status_details|JSON|
|target|String|
|zone|String|