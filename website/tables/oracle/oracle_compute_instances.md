# Table: oracle_compute_instances

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
|lifecycle_state|String|
|shape|String|
|time_created|Timestamp|
|capacity_reservation_id|String|
|dedicated_vm_host_id|String|
|defined_tags|JSON|
|display_name|String|
|extended_metadata|JSON|
|fault_domain|String|
|freeform_tags|JSON|
|image_id|String|
|ipxe_script|String|
|launch_mode|String|
|launch_options|JSON|
|instance_options|JSON|
|availability_config|JSON|
|preemptible_instance_config|JSON|
|metadata|JSON|
|shape_config|JSON|
|system_tags|JSON|
|agent_config|JSON|
|time_maintenance_reboot_due|Timestamp|