# Table: oracle_database_exadata_infrastructures

This table shows data for Oracle Database Exadata Infrastructures.

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
|lifecycle_state|`utf8`|
|display_name|`utf8`|
|shape|`utf8`|
|time_zone|`utf8`|
|cpus_enabled|`int64`|
|max_cpu_count|`int64`|
|memory_size_in_g_bs|`int64`|
|max_memory_in_g_bs|`int64`|
|db_node_storage_size_in_g_bs|`int64`|
|max_db_node_storage_in_g_bs|`int64`|
|data_storage_size_in_t_bs|`float64`|
|max_data_storage_in_t_bs|`float64`|
|rack_serial_number|`utf8`|
|storage_count|`int64`|
|additional_storage_count|`int64`|
|activated_storage_count|`int64`|
|compute_count|`int64`|
|is_multi_rack_deployment|`bool`|
|multi_rack_configuration_file|`binary`|
|additional_compute_count|`int64`|
|additional_compute_system_model|`utf8`|
|cloud_control_plane_server1|`utf8`|
|cloud_control_plane_server2|`utf8`|
|netmask|`utf8`|
|gateway|`utf8`|
|admin_network_cidr|`utf8`|
|infini_band_network_cidr|`utf8`|
|corporate_proxy|`utf8`|
|dns_server|`list<item: utf8, nullable>`|
|ntp_server|`list<item: utf8, nullable>`|
|time_created|`timestamp[us, tz=UTC]`|
|lifecycle_details|`utf8`|
|csi_number|`utf8`|
|contacts|`json`|
|maintenance_slo_status|`utf8`|
|maintenance_window|`json`|
|storage_server_version|`utf8`|
|db_server_version|`utf8`|
|monthly_db_server_version|`utf8`|
|last_maintenance_run_id|`utf8`|
|next_maintenance_run_id|`utf8`|
|is_cps_offline_report_enabled|`bool`|
|freeform_tags|`json`|
|defined_tags|`json`|