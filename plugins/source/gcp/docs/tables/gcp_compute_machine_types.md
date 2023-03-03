# Table: gcp_compute_machine_types

https://cloud.google.com/compute/docs/reference/rest/v1/machineTypes/list#response-body

The primary key for this table is **self_link**.

## Relations

This table depends on [gcp_compute_zones](gcp_compute_zones.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|accelerators|JSON|
|creation_timestamp|String|
|deprecated|JSON|
|description|String|
|guest_cpus|Int|
|id|Int|
|image_space_gb|Int|
|is_shared_cpu|Bool|
|kind|String|
|maximum_persistent_disks|Int|
|maximum_persistent_disks_size_gb|Int|
|memory_mb|Int|
|name|String|
|scratch_disks|JSON|
|self_link (PK)|String|
|zone|String|