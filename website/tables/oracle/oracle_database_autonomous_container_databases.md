# Table: oracle_database_autonomous_container_databases

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
|display_name|String|
|service_level_agreement_type|String|
|lifecycle_state|String|
|patch_model|String|
|db_unique_name|String|
|autonomous_exadata_infrastructure_id|String|
|autonomous_vm_cluster_id|String|
|infrastructure_type|String|
|cloud_autonomous_vm_cluster_id|String|
|kms_key_id|String|
|vault_id|String|
|kms_key_version_id|String|
|key_history_entry|JSON|
|lifecycle_details|String|
|time_created|Timestamp|
|patch_id|String|
|last_maintenance_run_id|String|
|next_maintenance_run_id|String|
|maintenance_window|JSON|
|standby_maintenance_buffer_in_days|Int|
|freeform_tags|JSON|
|defined_tags|JSON|
|role|String|
|availability_domain|String|
|db_version|String|
|backup_config|JSON|
|key_store_id|String|
|key_store_wallet_name|String|
|memory_per_oracle_compute_unit_in_g_bs|Int|
|available_cpus|Float|
|total_cpus|Int|
|reclaimable_cpus|Float|
|provisionable_cpus|JSON|
|compute_model|String|