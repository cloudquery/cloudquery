# Table: oracle_compute_dedicated_vm_hosts

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|region (PK)|String|
|compartment_id (PK)|String|
|id (PK)|String|
|availability_domain|String|
|dedicated_vm_host_shape|String|
|display_name|String|
|lifecycle_state|String|
|time_created|Timestamp|
|remaining_ocpus|Float|
|total_ocpus|Float|
|fault_domain|String|
|total_memory_in_g_bs|Float|
|remaining_memory_in_g_bs|Float|