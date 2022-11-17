# Table: azure_sql_databases

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql#Database

The primary key for this table is **id**.

## Relations
This table depends on [azure_sql_servers](azure_sql_servers.md).

The following tables depend on azure_sql_databases:
  - [azure_sql_backup_long_term_retention_policies](azure_sql_backup_long_term_retention_policies.md)
  - [azure_sql_database_blob_auditing_policies](azure_sql_database_blob_auditing_policies.md)
  - [azure_sql_database_threat_detection_policies](azure_sql_database_threat_detection_policies.md)
  - [azure_sql_database_vulnerability_assessment_scans](azure_sql_database_vulnerability_assessment_scans.md)
  - [azure_sql_database_vulnerability_assessments](azure_sql_database_vulnerability_assessments.md)
  - [azure_sql_transparent_data_encryptions](azure_sql_transparent_data_encryptions.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|identity|JSON|
|auto_pause_delay|Int|
|catalog_collation|String|
|collation|String|
|create_mode|String|
|elastic_pool_id|String|
|federated_client_id|String|
|high_availability_replica_count|Int|
|is_ledger_on|Bool|
|license_type|String|
|long_term_retention_backup_resource_id|String|
|maintenance_configuration_id|String|
|max_size_bytes|Int|
|min_capacity|Float|
|read_scale|String|
|recoverable_database_id|String|
|recovery_services_recovery_point_id|String|
|requested_backup_storage_redundancy|String|
|restorable_dropped_database_id|String|
|restore_point_in_time|Timestamp|
|sample_name|String|
|secondary_type|String|
|source_database_deletion_date|Timestamp|
|source_database_id|String|
|source_resource_id|String|
|zone_redundant|Bool|
|creation_date|Timestamp|
|current_backup_storage_redundancy|String|
|current_sku|JSON|
|current_service_objective_name|String|
|database_id|String|
|default_secondary_location|String|
|earliest_restore_date|Timestamp|
|failover_group_id|String|
|is_infra_encryption_enabled|Bool|
|max_log_size_bytes|Int|
|paused_date|Timestamp|
|requested_service_objective_name|String|
|resumed_date|Timestamp|
|status|String|
|sku|JSON|
|tags|JSON|
|id (PK)|String|
|kind|String|
|managed_by|String|
|name|String|
|type|String|
|server_id|String|