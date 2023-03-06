# Table: oracle_database_cloud_vm_clusters

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
|shape|String|
|lifecycle_state|String|
|display_name|String|
|hostname|String|
|domain|String|
|cpu_core_count|Int|
|cloud_exadata_infrastructure_id|String|
|ssh_public_keys|StringArray|
|backup_subnet_id|String|
|nsg_ids|StringArray|
|backup_network_nsg_ids|StringArray|
|last_update_history_entry_id|String|
|listener_port|Int|
|node_count|Int|
|storage_size_in_g_bs|Int|
|time_created|Timestamp|
|lifecycle_details|String|
|time_zone|String|
|ocpu_count|Float|
|memory_size_in_g_bs|Int|
|db_node_storage_size_in_g_bs|Int|
|data_storage_size_in_t_bs|Float|
|db_servers|StringArray|
|cluster_name|String|
|data_storage_percentage|Int|
|is_local_backup_enabled|Bool|
|is_sparse_diskgroup_enabled|Bool|
|gi_version|String|
|system_version|String|
|license_model|String|
|disk_redundancy|String|
|scan_ip_ids|StringArray|
|vip_ids|StringArray|
|scan_dns_record_id|String|
|freeform_tags|JSON|
|defined_tags|JSON|
|scan_dns_name|String|
|zone_id|String|
|scan_listener_port_tcp|Int|
|scan_listener_port_tcp_ssl|Int|
|data_collection_options|JSON|