# Table: gcp_vmmigration_sources

https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.sources

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_vmmigration_sources:
  - [gcp_vmmigration_source_datacenter_connectors](gcp_vmmigration_source_datacenter_connectors.md)
  - [gcp_vmmigration_source_migrating_vms](gcp_vmmigration_source_migrating_vms.md)
  - [gcp_vmmigration_source_utilization_reports](gcp_vmmigration_source_utilization_reports.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|create_time|Timestamp|
|update_time|Timestamp|
|labels|JSON|
|description|String|