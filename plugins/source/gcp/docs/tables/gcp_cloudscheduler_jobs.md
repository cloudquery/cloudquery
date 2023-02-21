# Table: gcp_cloudscheduler_jobs

https://cloud.google.com/scheduler/docs/reference/rest/v1/projects.locations.jobs#Job

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_cloudscheduler_locations](gcp_cloudscheduler_locations.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|description|String|
|schedule|String|
|time_zone|String|
|user_update_time|Timestamp|
|state|String|
|status|JSON|
|schedule_time|Timestamp|
|last_attempt_time|Timestamp|
|retry_config|JSON|
|attempt_deadline|Int|