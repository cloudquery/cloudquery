# Table: oracle_compute_instances

This table shows data for Oracle Compute Instances.

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
|lifecycle_state|`utf8`|
|shape|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|capacity_reservation_id|`utf8`|
|dedicated_vm_host_id|`utf8`|
|defined_tags|`json`|
|display_name|`utf8`|
|extended_metadata|`json`|
|fault_domain|`utf8`|
|freeform_tags|`json`|
|image_id|`utf8`|
|ipxe_script|`utf8`|
|launch_mode|`utf8`|
|launch_options|`json`|
|instance_options|`json`|
|availability_config|`json`|
|preemptible_instance_config|`json`|
|metadata|`json`|
|shape_config|`json`|
|system_tags|`json`|
|agent_config|`json`|
|time_maintenance_reboot_due|`timestamp[us, tz=UTC]`|