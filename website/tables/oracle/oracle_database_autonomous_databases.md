# Table: oracle_database_autonomous_databases

This table shows data for Oracle Database Autonomous Databases.

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
|db_name|`utf8`|
|cpu_core_count|`int64`|
|data_storage_size_in_t_bs|`int64`|
|lifecycle_details|`utf8`|
|kms_key_id|`utf8`|
|vault_id|`utf8`|
|kms_key_lifecycle_details|`utf8`|
|kms_key_version_id|`utf8`|
|character_set|`utf8`|
|ncharacter_set|`utf8`|
|is_free_tier|`bool`|
|system_tags|`json`|
|time_reclamation_of_free_autonomous_database|`timestamp[us, tz=UTC]`|
|time_deletion_of_free_autonomous_database|`timestamp[us, tz=UTC]`|
|backup_config|`json`|
|key_history_entry|`json`|
|compute_model|`utf8`|
|compute_count|`float64`|
|ocpu_count|`float64`|
|provisionable_cpus|`list<item: float64, nullable>`|
|memory_per_oracle_compute_unit_in_g_bs|`int64`|
|data_storage_size_in_g_bs|`int64`|
|infrastructure_type|`utf8`|
|is_dedicated|`bool`|
|autonomous_container_database_id|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|display_name|`utf8`|
|service_console_url|`utf8`|
|connection_strings|`json`|
|connection_urls|`json`|
|license_model|`utf8`|
|used_data_storage_size_in_t_bs|`int64`|
|freeform_tags|`json`|
|defined_tags|`json`|
|subnet_id|`utf8`|
|nsg_ids|`list<item: utf8, nullable>`|
|private_endpoint|`utf8`|
|private_endpoint_label|`utf8`|
|private_endpoint_ip|`utf8`|
|db_version|`utf8`|
|is_preview|`bool`|
|db_workload|`utf8`|
|is_access_control_enabled|`bool`|
|whitelisted_ips|`list<item: utf8, nullable>`|
|are_primary_whitelisted_ips_used|`bool`|
|standby_whitelisted_ips|`list<item: utf8, nullable>`|
|apex_details|`json`|
|is_auto_scaling_enabled|`bool`|
|data_safe_status|`utf8`|
|operations_insights_status|`utf8`|
|database_management_status|`utf8`|
|time_maintenance_begin|`timestamp[us, tz=UTC]`|
|time_maintenance_end|`timestamp[us, tz=UTC]`|
|is_refreshable_clone|`bool`|
|time_of_last_refresh|`timestamp[us, tz=UTC]`|
|time_of_last_refresh_point|`timestamp[us, tz=UTC]`|
|time_of_next_refresh|`timestamp[us, tz=UTC]`|
|open_mode|`utf8`|
|refreshable_status|`utf8`|
|refreshable_mode|`utf8`|
|source_id|`utf8`|
|permission_level|`utf8`|
|time_of_last_switchover|`timestamp[us, tz=UTC]`|
|time_of_last_failover|`timestamp[us, tz=UTC]`|
|is_data_guard_enabled|`bool`|
|failed_data_recovery_in_seconds|`int64`|
|standby_db|`json`|
|is_local_data_guard_enabled|`bool`|
|is_remote_data_guard_enabled|`bool`|
|local_standby_db|`json`|
|role|`utf8`|
|available_upgrade_versions|`list<item: utf8, nullable>`|
|key_store_id|`utf8`|
|key_store_wallet_name|`utf8`|
|supported_regions_to_clone_to|`list<item: utf8, nullable>`|
|customer_contacts|`json`|
|time_local_data_guard_enabled|`timestamp[us, tz=UTC]`|
|dataguard_region_type|`utf8`|
|time_data_guard_role_changed|`timestamp[us, tz=UTC]`|
|peer_db_ids|`list<item: utf8, nullable>`|
|is_mtls_connection_required|`bool`|
|is_reconnect_clone_enabled|`bool`|
|time_until_reconnect_clone_enabled|`timestamp[us, tz=UTC]`|
|autonomous_maintenance_schedule_type|`utf8`|
|scheduled_operations|`json`|
|is_auto_scaling_for_storage_enabled|`bool`|
|allocated_storage_size_in_t_bs|`float64`|
|actual_used_data_storage_size_in_t_bs|`float64`|
|max_cpu_core_count|`int64`|
|database_edition|`utf8`|
|db_tools_details|`json`|