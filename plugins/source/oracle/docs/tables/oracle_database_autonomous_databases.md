# Table: oracle_database_autonomous_databases

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
|db_name|String|
|cpu_core_count|Int|
|data_storage_size_in_t_bs|Int|
|lifecycle_details|String|
|kms_key_id|String|
|vault_id|String|
|kms_key_lifecycle_details|String|
|kms_key_version_id|String|
|character_set|String|
|ncharacter_set|String|
|is_free_tier|Bool|
|system_tags|JSON|
|time_reclamation_of_free_autonomous_database|Timestamp|
|time_deletion_of_free_autonomous_database|Timestamp|
|backup_config|JSON|
|key_history_entry|JSON|
|ocpu_count|Float|
|provisionable_cpus|JSON|
|memory_per_oracle_compute_unit_in_g_bs|Int|
|data_storage_size_in_g_bs|Int|
|infrastructure_type|String|
|is_dedicated|Bool|
|autonomous_container_database_id|String|
|time_created|Timestamp|
|display_name|String|
|service_console_url|String|
|connection_strings|JSON|
|connection_urls|JSON|
|license_model|String|
|used_data_storage_size_in_t_bs|Int|
|freeform_tags|JSON|
|defined_tags|JSON|
|subnet_id|String|
|nsg_ids|StringArray|
|private_endpoint|String|
|private_endpoint_label|String|
|private_endpoint_ip|String|
|db_version|String|
|is_preview|Bool|
|db_workload|String|
|is_access_control_enabled|Bool|
|whitelisted_ips|StringArray|
|are_primary_whitelisted_ips_used|Bool|
|standby_whitelisted_ips|StringArray|
|apex_details|JSON|
|is_auto_scaling_enabled|Bool|
|data_safe_status|String|
|operations_insights_status|String|
|database_management_status|String|
|time_maintenance_begin|Timestamp|
|time_maintenance_end|Timestamp|
|is_refreshable_clone|Bool|
|time_of_last_refresh|Timestamp|
|time_of_last_refresh_point|Timestamp|
|time_of_next_refresh|Timestamp|
|open_mode|String|
|refreshable_status|String|
|refreshable_mode|String|
|source_id|String|
|permission_level|String|
|time_of_last_switchover|Timestamp|
|time_of_last_failover|Timestamp|
|is_data_guard_enabled|Bool|
|failed_data_recovery_in_seconds|Int|
|standby_db|JSON|
|is_local_data_guard_enabled|Bool|
|is_remote_data_guard_enabled|Bool|
|local_standby_db|JSON|
|role|String|
|available_upgrade_versions|StringArray|
|key_store_id|String|
|key_store_wallet_name|String|
|supported_regions_to_clone_to|StringArray|
|customer_contacts|JSON|
|time_local_data_guard_enabled|Timestamp|
|dataguard_region_type|String|
|time_data_guard_role_changed|Timestamp|
|peer_db_ids|StringArray|
|is_mtls_connection_required|Bool|
|is_reconnect_clone_enabled|Bool|
|time_until_reconnect_clone_enabled|Timestamp|
|autonomous_maintenance_schedule_type|String|
|scheduled_operations|JSON|
|is_auto_scaling_for_storage_enabled|Bool|
|allocated_storage_size_in_t_bs|Float|
|actual_used_data_storage_size_in_t_bs|Float|
|max_cpu_core_count|Int|
|database_edition|String|