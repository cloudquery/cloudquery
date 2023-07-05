# Table: gcp_vmmigration_source_migrating_vm_cutover_jobs

This table shows data for GCP VM Migration Source Migrating VM Cutover Jobs.

https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.sources.migratingVms.cutoverJobs

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_vmmigration_source_migrating_vms](gcp_vmmigration_source_migrating_vms).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|end_time|`timestamp[us, tz=UTC]`|
|name (PK)|`utf8`|
|state|`utf8`|
|state_time|`timestamp[us, tz=UTC]`|
|progress_percent|`int64`|
|error|`json`|
|state_message|`utf8`|
|steps|`json`|