# Table: oracle_compute_dedicated_vm_hosts

This table shows data for Oracle Compute Dedicated Virtual Machine (VM) Hosts.

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|id (PK)|`utf8`|
|availability_domain|`utf8`|
|dedicated_vm_host_shape|`utf8`|
|display_name|`utf8`|
|lifecycle_state|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|remaining_ocpus|`float64`|
|total_ocpus|`float64`|
|fault_domain|`utf8`|
|total_memory_in_g_bs|`float64`|
|remaining_memory_in_g_bs|`float64`|