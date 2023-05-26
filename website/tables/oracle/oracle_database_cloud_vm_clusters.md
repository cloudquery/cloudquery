# Table: oracle_database_cloud_vm_clusters

This table shows data for Oracle Database Cloud Virtual Machine (VM) Clusters.

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
|shape|`utf8`|
|lifecycle_state|`utf8`|
|display_name|`utf8`|
|hostname|`utf8`|
|domain|`utf8`|
|cpu_core_count|`int64`|
|cloud_exadata_infrastructure_id|`utf8`|
|ssh_public_keys|`list<item: utf8, nullable>`|
|backup_subnet_id|`utf8`|
|nsg_ids|`list<item: utf8, nullable>`|
|backup_network_nsg_ids|`list<item: utf8, nullable>`|
|last_update_history_entry_id|`utf8`|
|listener_port|`int64`|
|node_count|`int64`|
|storage_size_in_g_bs|`int64`|
|time_created|`timestamp[us, tz=UTC]`|
|lifecycle_details|`utf8`|
|time_zone|`utf8`|
|ocpu_count|`float64`|
|memory_size_in_g_bs|`int64`|
|db_node_storage_size_in_g_bs|`int64`|
|data_storage_size_in_t_bs|`float64`|
|db_servers|`list<item: utf8, nullable>`|
|cluster_name|`utf8`|
|data_storage_percentage|`int64`|
|is_local_backup_enabled|`bool`|
|is_sparse_diskgroup_enabled|`bool`|
|gi_version|`utf8`|
|system_version|`utf8`|
|license_model|`utf8`|
|disk_redundancy|`utf8`|
|scan_ip_ids|`list<item: utf8, nullable>`|
|vip_ids|`list<item: utf8, nullable>`|
|scan_dns_record_id|`utf8`|
|freeform_tags|`json`|
|defined_tags|`json`|
|scan_dns_name|`utf8`|
|zone_id|`utf8`|
|scan_listener_port_tcp|`int64`|
|scan_listener_port_tcp_ssl|`int64`|
|data_collection_options|`json`|