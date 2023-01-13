# Table: gcp_vmmigration_source_datacenter_connectors

https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.sources.datacenterConnectors

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
|create_time|Timestamp|
|update_time|Timestamp|
|registration_id|String|
|service_account|String|
|version|String|
|bucket|String|
|state|String|
|state_time|Timestamp|
|error|JSON|
|appliance_infrastructure_version|String|
|appliance_software_version|String|
|available_versions|JSON|
|upgrade_status|JSON|