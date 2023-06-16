# Table: oracle_database_autonomous_exadata_infrastructures

This table shows data for Oracle Database Autonomous Exadata Infrastructures.

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
|display_name|`utf8`|
|availability_domain|`utf8`|
|subnet_id|`utf8`|
|shape|`utf8`|
|hostname|`utf8`|
|domain|`utf8`|
|lifecycle_state|`utf8`|
|maintenance_window|`json`|
|nsg_ids|`list<item: utf8, nullable>`|
|lifecycle_details|`utf8`|
|license_model|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|last_maintenance_run_id|`utf8`|
|next_maintenance_run_id|`utf8`|
|freeform_tags|`json`|
|defined_tags|`json`|
|scan_dns_name|`utf8`|
|zone_id|`utf8`|