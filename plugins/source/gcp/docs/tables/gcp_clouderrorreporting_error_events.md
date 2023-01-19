# Table: gcp_clouderrorreporting_error_events

https://cloud.google.com/error-reporting/reference/rest/v1beta1/ErrorEvent

The primary key for this table is **_cq_id**.

## Relations

This table depends on [gcp_clouderrorreporting_error_group_stats](gcp_clouderrorreporting_error_group_stats.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|event_time|Timestamp|
|service_context|JSON|
|message|String|
|context|JSON|