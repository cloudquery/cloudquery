# Table: oracle_database_cloud_exadata_infrastructures

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
|lifecycle_state|String|
|display_name|String|
|shape|String|
|availability_domain|String|
|compute_count|Int|
|storage_count|Int|
|total_storage_size_in_g_bs|Int|
|available_storage_size_in_g_bs|Int|
|cpu_count|Int|
|max_cpu_count|Int|
|memory_size_in_g_bs|Int|
|max_memory_in_g_bs|Int|
|db_node_storage_size_in_g_bs|Int|
|max_db_node_storage_in_g_bs|Int|
|data_storage_size_in_t_bs|Float|
|max_data_storage_in_t_bs|Float|
|additional_storage_count|Int|
|activated_storage_count|Int|
|time_created|Timestamp|
|lifecycle_details|String|
|maintenance_window|JSON|
|last_maintenance_run_id|String|
|next_maintenance_run_id|String|
|freeform_tags|JSON|
|defined_tags|JSON|
|customer_contacts|JSON|