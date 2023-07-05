# Table: gcp_vmmigration_source_migrating_vms

This table shows data for GCP VM Migration Source Migrating Virtual Machines (VMs).

https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.sources.migratingVms

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_vmmigration_sources](gcp_vmmigration_sources).

The following tables depend on gcp_vmmigration_source_migrating_vms:
  - [gcp_vmmigration_source_migrating_vm_clone_jobs](gcp_vmmigration_source_migrating_vm_clone_jobs)
  - [gcp_vmmigration_source_migrating_vm_cutover_jobs](gcp_vmmigration_source_migrating_vm_cutover_jobs)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|source_vm_id|`utf8`|
|display_name|`utf8`|
|description|`utf8`|
|policy|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|last_sync|`json`|
|state|`utf8`|
|state_time|`timestamp[us, tz=UTC]`|
|current_sync_info|`json`|
|group|`utf8`|
|labels|`json`|
|recent_clone_jobs|`json`|
|error|`json`|
|recent_cutover_jobs|`json`|