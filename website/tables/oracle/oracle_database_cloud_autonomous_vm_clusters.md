# Table: oracle_database_cloud_autonomous_vm_clusters

This table shows data for Oracle Database Cloud Autonomous Virtual Machine (VM) Clusters.

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
|subnet_id|`utf8`|
|lifecycle_state|`utf8`|
|display_name|`utf8`|
|cloud_exadata_infrastructure_id|`utf8`|
|description|`utf8`|
|nsg_ids|`list<item: utf8, nullable>`|
|last_update_history_entry_id|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|time_updated|`timestamp[us, tz=UTC]`|
|cluster_time_zone|`utf8`|
|lifecycle_details|`utf8`|
|hostname|`utf8`|
|domain|`utf8`|
|shape|`utf8`|
|node_count|`int64`|
|data_storage_size_in_t_bs|`float64`|
|data_storage_size_in_g_bs|`float64`|
|cpu_core_count|`int64`|
|ocpu_count|`float64`|
|compute_model|`utf8`|
|cpu_core_count_per_node|`int64`|
|memory_size_in_g_bs|`int64`|
|license_model|`utf8`|
|last_maintenance_run_id|`utf8`|
|next_maintenance_run_id|`utf8`|
|maintenance_window|`json`|
|freeform_tags|`json`|
|defined_tags|`json`|
|available_cpus|`float64`|
|reclaimable_cpus|`float64`|
|available_container_databases|`int64`|
|total_container_databases|`int64`|
|available_autonomous_data_storage_size_in_t_bs|`float64`|
|autonomous_data_storage_size_in_t_bs|`float64`|
|db_node_storage_size_in_g_bs|`int64`|
|memory_per_oracle_compute_unit_in_g_bs|`int64`|