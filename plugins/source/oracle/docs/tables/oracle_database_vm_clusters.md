# Table: oracle_database_vm_clusters

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
|last_patch_history_entry_id|String|
|lifecycle_state|String|
|display_name|String|
|time_created|Timestamp|
|lifecycle_details|String|
|time_zone|String|
|is_local_backup_enabled|Bool|
|exadata_infrastructure_id|String|
|is_sparse_diskgroup_enabled|Bool|
|vm_cluster_network_id|String|
|cpus_enabled|Int|
|ocpus_enabled|Float|
|memory_size_in_g_bs|Int|
|db_node_storage_size_in_g_bs|Int|
|data_storage_size_in_t_bs|Float|
|data_storage_size_in_g_bs|Float|
|shape|String|
|gi_version|String|
|system_version|String|
|ssh_public_keys|StringArray|
|license_model|String|
|db_servers|StringArray|
|freeform_tags|JSON|
|defined_tags|JSON|
|data_collection_options|JSON|