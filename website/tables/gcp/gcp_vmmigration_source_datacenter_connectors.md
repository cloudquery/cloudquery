# Table: gcp_vmmigration_source_datacenter_connectors

This table shows data for GCP VM Migration Source Datacenter Connectors.

https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.sources.datacenterConnectors

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_vmmigration_sources](gcp_vmmigration_sources).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|name (PK)|`utf8`|
|registration_id|`utf8`|
|service_account|`utf8`|
|version|`utf8`|
|bucket|`utf8`|
|state|`utf8`|
|state_time|`timestamp[us, tz=UTC]`|
|error|`json`|
|appliance_infrastructure_version|`utf8`|
|appliance_software_version|`utf8`|
|available_versions|`json`|
|upgrade_status|`json`|