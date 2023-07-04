# Table: gcp_vmmigration_sources

This table shows data for GCP VM Migration Sources.

https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.sources

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_vmmigration_sources:
  - [gcp_vmmigration_source_datacenter_connectors](gcp_vmmigration_source_datacenter_connectors)
  - [gcp_vmmigration_source_migrating_vms](gcp_vmmigration_source_migrating_vms)
  - [gcp_vmmigration_source_utilization_reports](gcp_vmmigration_source_utilization_reports)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|labels|`json`|
|description|`utf8`|