# Table: azure_sql_databases

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql#Database

The primary key for this table is **id**.

## Relations
This table depends on [azure_sql_servers](azure_sql_servers.md).

The following tables depend on azure_sql_databases:
  - [azure_sql_database_blob_auditing_policies](azure_sql_database_blob_auditing_policies.md)
  - [azure_sql_database_vulnerability_assessments](azure_sql_database_vulnerability_assessments.md)
  - [azure_sql_database_vulnerability_assessment_scans](azure_sql_database_vulnerability_assessment_scans.md)
  - [azure_sql_backup_long_term_retention_policies](azure_sql_backup_long_term_retention_policies.md)
  - [azure_sql_database_threat_detection_policies](azure_sql_database_threat_detection_policies.md)
  - [azure_sql_transparent_data_encryptions](azure_sql_transparent_data_encryptions.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|sql_server_id|String|
|sku|JSON|
|kind|String|
|managed_by|String|
|create_mode|String|
|collation|String|
|max_size_bytes|Int|
|sample_name|String|
|elastic_pool_id|String|
|source_database_id|String|
|status|String|
|database_id|UUID|
|creation_date|Timestamp|
|current_service_objective_name|String|
|requested_service_objective_name|String|
|default_secondary_location|String|
|failover_group_id|String|
|restore_point_in_time|Timestamp|
|source_database_deletion_date|Timestamp|
|recovery_services_recovery_point_id|String|
|long_term_retention_backup_resource_id|String|
|recoverable_database_id|String|
|restorable_dropped_database_id|String|
|catalog_collation|String|
|zone_redundant|Bool|
|license_type|String|
|max_log_size_bytes|Int|
|earliest_restore_date|Timestamp|
|read_scale|String|
|high_availability_replica_count|Int|
|secondary_type|String|
|current_sku|JSON|
|auto_pause_delay|Int|
|storage_account_type|String|
|min_capacity|Float|
|paused_date|Timestamp|
|resumed_date|Timestamp|
|maintenance_configuration_id|String|
|location|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|