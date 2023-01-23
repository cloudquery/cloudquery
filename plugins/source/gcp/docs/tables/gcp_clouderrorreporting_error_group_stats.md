# Table: gcp_clouderrorreporting_error_group_stats

https://cloud.google.com/error-reporting/reference/rest/v1beta1/projects.groupStats/list#ErrorGroupStats

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on gcp_clouderrorreporting_error_group_stats:
  - [gcp_clouderrorreporting_error_events](gcp_clouderrorreporting_error_events.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|group|JSON|
|count|Int|
|affected_users_count|Int|
|timed_counts|JSON|
|first_seen_time|Timestamp|
|last_seen_time|Timestamp|
|affected_services|JSON|
|num_affected_services|Int|
|representative|JSON|