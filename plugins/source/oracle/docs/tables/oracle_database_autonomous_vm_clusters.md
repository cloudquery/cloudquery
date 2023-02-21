# Table: oracle_database_autonomous_vm_clusters

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
|lifecycle_state|String|
|exadata_infrastructure_id|String|
|vm_cluster_network_id|String|
|time_created|Timestamp|
|lifecycle_details|String|
|time_zone|String|
|is_local_backup_enabled|Bool|
|cpus_enabled|Int|
|compute_model|String|
|ocpus_enabled|Float|
|available_cpus|Int|
|total_container_databases|Int|
|memory_per_oracle_compute_unit_in_g_bs|Int|
|cpu_core_count_per_node|Int|
|autonomous_data_storage_size_in_t_bs|Float|
|maintenance_window|JSON|
|last_maintenance_run_id|String|
|next_maintenance_run_id|String|
|memory_size_in_g_bs|Int|
|db_node_storage_size_in_g_bs|Int|
|data_storage_size_in_t_bs|Float|
|data_storage_size_in_g_bs|Float|
|available_data_storage_size_in_t_bs|Float|
|license_model|String|
|freeform_tags|JSON|
|defined_tags|JSON|
|reclaimable_cpus|Int|
|available_container_databases|Int|
|available_autonomous_data_storage_size_in_t_bs|Float|
|scan_listener_port_tls|Int|
|scan_listener_port_non_tls|Int|
|is_mtls_enabled|Bool|