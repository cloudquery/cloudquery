# Table: gcp_cloudscheduler_jobs

This table shows data for GCP Cloud Scheduler Jobs.

https://cloud.google.com/scheduler/docs/reference/rest/v1/projects.locations.jobs#Job

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_cloudscheduler_locations](gcp_cloudscheduler_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|description|`utf8`|
|schedule|`utf8`|
|time_zone|`utf8`|
|user_update_time|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|status|`json`|
|schedule_time|`timestamp[us, tz=UTC]`|
|last_attempt_time|`timestamp[us, tz=UTC]`|
|retry_config|`json`|
|attempt_deadline|`int64`|