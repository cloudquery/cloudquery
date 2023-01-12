# Table: oracle_database_cloud_autonomous_vm_clusters

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
|subnet_id|String|
|lifecycle_state|String|
|display_name|String|
|cloud_exadata_infrastructure_id|String|
|description|String|
|nsg_ids|StringArray|
|last_update_history_entry_id|String|
|time_created|Timestamp|
|time_updated|Timestamp|
|cluster_time_zone|String|
|lifecycle_details|String|
|hostname|String|
|domain|String|
|shape|String|
|node_count|Int|
|data_storage_size_in_t_bs|Float|
|data_storage_size_in_g_bs|Float|
|cpu_core_count|Int|
|ocpu_count|Float|
|cpu_core_count_per_node|Int|
|memory_size_in_g_bs|Int|
|license_model|String|
|last_maintenance_run_id|String|
|next_maintenance_run_id|String|
|maintenance_window|JSON|
|freeform_tags|JSON|
|defined_tags|JSON|
|available_cpus|Float|
|reclaimable_cpus|Float|
|available_container_databases|Int|
|total_container_databases|Int|
|available_autonomous_data_storage_size_in_t_bs|Float|
|autonomous_data_storage_size_in_t_bs|Float|
|db_node_storage_size_in_g_bs|Int|
|memory_per_oracle_compute_unit_in_g_bs|Int|