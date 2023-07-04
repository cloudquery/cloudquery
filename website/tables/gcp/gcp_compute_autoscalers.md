# Table: gcp_compute_autoscalers

This table shows data for GCP Compute Autoscalers.

https://cloud.google.com/compute/docs/reference/rest/v1/autoscalers#Autoscaler

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|autoscaling_policy|`json`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|id|`int64`|
|kind|`utf8`|
|name|`utf8`|
|recommended_size|`int64`|
|region|`utf8`|
|scaling_schedule_status|`json`|
|self_link (PK)|`utf8`|
|status|`utf8`|
|status_details|`json`|
|target|`utf8`|
|zone|`utf8`|