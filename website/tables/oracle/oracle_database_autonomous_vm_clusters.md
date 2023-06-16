# Table: oracle_database_autonomous_vm_clusters

This table shows data for Oracle Database Autonomous Virtual Machine (VM) Clusters.

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
|lifecycle_state|`utf8`|
|exadata_infrastructure_id|`utf8`|
|vm_cluster_network_id|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|lifecycle_details|`utf8`|
|time_zone|`utf8`|
|is_local_backup_enabled|`bool`|
|cpus_enabled|`int64`|
|compute_model|`utf8`|
|ocpus_enabled|`float64`|
|available_cpus|`int64`|
|total_container_databases|`int64`|
|memory_per_oracle_compute_unit_in_g_bs|`int64`|
|cpu_core_count_per_node|`int64`|
|autonomous_data_storage_size_in_t_bs|`float64`|
|maintenance_window|`json`|
|last_maintenance_run_id|`utf8`|
|next_maintenance_run_id|`utf8`|
|memory_size_in_g_bs|`int64`|
|db_node_storage_size_in_g_bs|`int64`|
|data_storage_size_in_t_bs|`float64`|
|data_storage_size_in_g_bs|`float64`|
|available_data_storage_size_in_t_bs|`float64`|
|license_model|`utf8`|
|freeform_tags|`json`|
|defined_tags|`json`|
|reclaimable_cpus|`int64`|
|available_container_databases|`int64`|
|available_autonomous_data_storage_size_in_t_bs|`float64`|
|scan_listener_port_tls|`int64`|
|scan_listener_port_non_tls|`int64`|
|is_mtls_enabled|`bool`|