# Table: gcp_clouderrorreporting_error_group_stats

This table shows data for GCP Clouderrorreporting Error Group Stats.

https://cloud.google.com/error-reporting/reference/rest/v1beta1/projects.groupStats/list#ErrorGroupStats

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on gcp_clouderrorreporting_error_group_stats:
  - [gcp_clouderrorreporting_error_events](gcp_clouderrorreporting_error_events)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|group|`json`|
|count|`int64`|
|affected_users_count|`int64`|
|timed_counts|`json`|
|first_seen_time|`timestamp[us, tz=UTC]`|
|last_seen_time|`timestamp[us, tz=UTC]`|
|affected_services|`json`|
|num_affected_services|`int64`|
|representative|`json`|