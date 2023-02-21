# Table: gcp_cloudscheduler_locations

https://cloud.google.com/scheduler/docs/reference/rest/v1/projects.locations#Location

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_cloudscheduler_locations:
  - [gcp_cloudscheduler_jobs](gcp_cloudscheduler_jobs.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|display_name|String|
|labels|JSON|
|location_id|String|
|metadata|ByteArray|
|name (PK)|String|