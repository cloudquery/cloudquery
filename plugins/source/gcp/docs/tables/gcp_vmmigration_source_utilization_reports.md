# Table: gcp_vmmigration_source_utilization_reports

https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.sources.utilizationReports

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_vmmigration_sources](gcp_vmmigration_sources.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|display_name|String|
|state|String|
|state_time|Timestamp|
|error|JSON|
|create_time|Timestamp|
|time_frame|String|
|frame_end_time|Timestamp|
|vm_count|Int|
|vms|JSON|