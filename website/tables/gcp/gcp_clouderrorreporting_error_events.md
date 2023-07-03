# Table: gcp_clouderrorreporting_error_events

This table shows data for GCP Clouderrorreporting Error Events.

https://cloud.google.com/error-reporting/reference/rest/v1beta1/ErrorEvent

The composite primary key for this table is ().

## Relations

This table depends on [gcp_clouderrorreporting_error_group_stats](gcp_clouderrorreporting_error_group_stats).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|project_id|`utf8`|
|event_time|`timestamp[us, tz=UTC]`|
|service_context|`json`|
|message|`utf8`|
|context|`json`|