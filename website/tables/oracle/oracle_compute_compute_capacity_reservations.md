# Table: oracle_compute_compute_capacity_reservations

This table shows data for Oracle Compute Compute Capacity Reservations.

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
|time_created|Timestamp|
|display_name|String|
|defined_tags|JSON|
|freeform_tags|JSON|
|lifecycle_state|String|
|reserved_instance_count|Int|
|used_instance_count|Int|
|is_default_reservation|Bool|