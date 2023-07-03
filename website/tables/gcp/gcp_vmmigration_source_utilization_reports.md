# Table: gcp_vmmigration_source_utilization_reports

This table shows data for GCP VM Migration Source Utilization Reports.

https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.sources.utilizationReports

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_vmmigration_sources](gcp_vmmigration_sources).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|state|`utf8`|
|state_time|`timestamp[us, tz=UTC]`|
|error|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|time_frame|`utf8`|
|frame_end_time|`timestamp[us, tz=UTC]`|
|vm_count|`int64`|
|vms|`json`|