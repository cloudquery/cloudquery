# Table: gcp_vmmigration_source_migrating_vms

https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.sources.migratingVms

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_vmmigration_sources](gcp_vmmigration_sources.md).

The following tables depend on gcp_vmmigration_source_migrating_vms:
  - [gcp_vmmigration_source_migrating_vm_clone_jobs](gcp_vmmigration_source_migrating_vm_clone_jobs.md)
  - [gcp_vmmigration_source_migrating_vm_cutover_jobs](gcp_vmmigration_source_migrating_vm_cutover_jobs.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|source_vm_id|String|
|display_name|String|
|description|String|
|policy|JSON|
|create_time|Timestamp|
|update_time|Timestamp|
|last_sync|JSON|
|state|String|
|state_time|Timestamp|
|current_sync_info|JSON|
|group|String|
|labels|JSON|
|recent_clone_jobs|JSON|
|error|JSON|
|recent_cutover_jobs|JSON|