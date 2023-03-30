# Table: gcp_vmmigration_source_migrating_vm_cutover_jobs

This table shows data for GCP VM Migration Source Migrating VM Cutover Jobs.

https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.sources.migratingVms.cutoverJobs

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_vmmigration_source_migrating_vms](gcp_vmmigration_source_migrating_vms).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|create_time|Timestamp|
|end_time|Timestamp|
|name (PK)|String|
|state|String|
|state_time|Timestamp|
|progress_percent|Int|
|error|JSON|
|state_message|String|
|steps|JSON|