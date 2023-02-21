# Table: oracle_database_autonomous_exadata_infrastructures

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
|display_name|String|
|availability_domain|String|
|subnet_id|String|
|shape|String|
|hostname|String|
|domain|String|
|lifecycle_state|String|
|maintenance_window|JSON|
|nsg_ids|StringArray|
|lifecycle_details|String|
|license_model|String|
|time_created|Timestamp|
|last_maintenance_run_id|String|
|next_maintenance_run_id|String|
|freeform_tags|JSON|
|defined_tags|JSON|
|scan_dns_name|String|
|zone_id|String|