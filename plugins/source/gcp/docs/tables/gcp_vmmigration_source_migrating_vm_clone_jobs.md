# Table: gcp_vmmigration_source_migrating_vm_clone_jobs

https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.sources.migratingVms.cloneJobs

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_vmmigration_source_migrating_vms](gcp_vmmigration_source_migrating_vms.md).

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
|error|JSON|