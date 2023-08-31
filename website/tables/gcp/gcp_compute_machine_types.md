# Table: gcp_compute_machine_types

This table shows data for GCP Compute Machine Types.

https://cloud.google.com/compute/docs/reference/rest/v1/machineTypes/list#response-body

The primary key for this table is **self_link**.

## Relations

This table depends on [gcp_compute_zones](gcp_compute_zones).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|accelerators|`json`|
|creation_timestamp|`utf8`|
|deprecated|`json`|
|description|`utf8`|
|guest_cpus|`int64`|
|id|`int64`|
|image_space_gb|`int64`|
|is_shared_cpu|`bool`|
|kind|`utf8`|
|maximum_persistent_disks|`int64`|
|maximum_persistent_disks_size_gb|`int64`|
|memory_mb|`int64`|
|name|`utf8`|
|scratch_disks|`json`|
|self_link (PK)|`utf8`|
|zone|`utf8`|