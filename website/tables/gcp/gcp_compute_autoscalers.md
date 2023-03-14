# Table: gcp_compute_autoscalers

This table shows data for GCP Compute Autoscalers.

https://cloud.google.com/compute/docs/reference/rest/v1/autoscalers#Autoscaler

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|autoscaling_policy|JSON|
|creation_timestamp|String|
|description|String|
|id|Int|
|kind|String|
|name|String|
|recommended_size|Int|
|region|String|
|scaling_schedule_status|JSON|
|self_link (PK)|String|
|status|String|
|status_details|JSON|
|target|String|
|zone|String|