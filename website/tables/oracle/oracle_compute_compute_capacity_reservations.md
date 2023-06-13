# Table: oracle_compute_compute_capacity_reservations

This table shows data for Oracle Compute Compute Capacity Reservations.

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
|time_created|`timestamp[us, tz=UTC]`|
|display_name|`utf8`|
|defined_tags|`json`|
|freeform_tags|`json`|
|lifecycle_state|`utf8`|
|reserved_instance_count|`int64`|
|used_instance_count|`int64`|
|is_default_reservation|`bool`|