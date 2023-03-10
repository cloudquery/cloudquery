# Table: oracle_database_exadata_infrastructures

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
|time_zone|String|
|cpus_enabled|Int|
|max_cpu_count|Int|
|memory_size_in_g_bs|Int|
|max_memory_in_g_bs|Int|
|db_node_storage_size_in_g_bs|Int|
|max_db_node_storage_in_g_bs|Int|
|data_storage_size_in_t_bs|Float|
|max_data_storage_in_t_bs|Float|
|rack_serial_number|String|
|storage_count|Int|
|additional_storage_count|Int|
|activated_storage_count|Int|
|compute_count|Int|
|is_multi_rack_deployment|Bool|
|multi_rack_configuration_file|ByteArray|
|additional_compute_count|Int|
|additional_compute_system_model|String|
|cloud_control_plane_server1|String|
|cloud_control_plane_server2|String|
|netmask|String|
|gateway|String|
|admin_network_cidr|String|
|infini_band_network_cidr|String|
|corporate_proxy|String|
|dns_server|StringArray|
|ntp_server|StringArray|
|time_created|Timestamp|
|lifecycle_details|String|
|csi_number|String|
|contacts|JSON|
|maintenance_slo_status|String|
|maintenance_window|JSON|
|storage_server_version|String|
|db_server_version|String|
|monthly_db_server_version|String|
|last_maintenance_run_id|String|
|next_maintenance_run_id|String|
|is_cps_offline_report_enabled|Bool|
|freeform_tags|JSON|
|defined_tags|JSON|