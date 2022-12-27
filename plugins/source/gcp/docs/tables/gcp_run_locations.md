# Table: gcp_run_locations

https://cloud.google.com/run/docs/reference/rest/v1/projects.locations#Location

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on gcp_run_locations:
  - [gcp_run_services](gcp_run_services.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|display_name|String|
|labels|JSON|
|location_id|String|
|metadata|IntArray|
|name|String|