# Table: oracle_database_vm_clusters

This table shows data for Oracle Database Virtual Machine (VM) Clusters.

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
|last_patch_history_entry_id|`utf8`|
|lifecycle_state|`utf8`|
|display_name|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|lifecycle_details|`utf8`|
|time_zone|`utf8`|
|is_local_backup_enabled|`bool`|
|exadata_infrastructure_id|`utf8`|
|is_sparse_diskgroup_enabled|`bool`|
|vm_cluster_network_id|`utf8`|
|cpus_enabled|`int64`|
|ocpus_enabled|`float64`|
|memory_size_in_g_bs|`int64`|
|db_node_storage_size_in_g_bs|`int64`|
|data_storage_size_in_t_bs|`float64`|
|data_storage_size_in_g_bs|`float64`|
|shape|`utf8`|
|gi_version|`utf8`|
|system_version|`utf8`|
|ssh_public_keys|`list<item: utf8, nullable>`|
|license_model|`utf8`|
|db_servers|`list<item: utf8, nullable>`|
|freeform_tags|`json`|
|defined_tags|`json`|
|data_collection_options|`json`|