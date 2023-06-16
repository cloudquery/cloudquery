# Table: oracle_database_autonomous_container_databases

This table shows data for Oracle Database Autonomous Container Databases.

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
|service_level_agreement_type|`utf8`|
|lifecycle_state|`utf8`|
|patch_model|`utf8`|
|db_unique_name|`utf8`|
|autonomous_exadata_infrastructure_id|`utf8`|
|autonomous_vm_cluster_id|`utf8`|
|infrastructure_type|`utf8`|
|cloud_autonomous_vm_cluster_id|`utf8`|
|kms_key_id|`utf8`|
|vault_id|`utf8`|
|kms_key_version_id|`utf8`|
|key_history_entry|`json`|
|lifecycle_details|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|patch_id|`utf8`|
|last_maintenance_run_id|`utf8`|
|next_maintenance_run_id|`utf8`|
|maintenance_window|`json`|
|standby_maintenance_buffer_in_days|`int64`|
|freeform_tags|`json`|
|defined_tags|`json`|
|role|`utf8`|
|availability_domain|`utf8`|
|db_version|`utf8`|
|backup_config|`json`|
|key_store_id|`utf8`|
|key_store_wallet_name|`utf8`|
|memory_per_oracle_compute_unit_in_g_bs|`int64`|
|available_cpus|`float64`|
|total_cpus|`int64`|
|reclaimable_cpus|`float64`|
|provisionable_cpus|`list<item: float64, nullable>`|
|compute_model|`utf8`|