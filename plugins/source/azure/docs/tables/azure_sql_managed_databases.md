# Table: azure_sql_managed_databases

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql#ManagedDatabase

The primary key for this table is **id**.

## Relations
This table depends on [azure_sql_managed_instances](azure_sql_managed_instances.md).

The following tables depend on azure_sql_managed_databases:
  - [azure_sql_managed_database_vulnerability_assessments](azure_sql_managed_database_vulnerability_assessments.md)
  - [azure_sql_managed_database_vulnerability_assessment_scans](azure_sql_managed_database_vulnerability_assessment_scans.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|sql_managed_instance_id|String|
|collation|String|
|status|String|
|creation_date|Timestamp|
|earliest_restore_point|Timestamp|
|restore_point_in_time|Timestamp|
|default_secondary_location|String|
|catalog_collation|String|
|create_mode|String|
|storage_container_uri|String|
|source_database_id|String|
|restorable_dropped_database_id|String|
|storage_container_sas_token|String|
|failover_group_id|String|
|recoverable_database_id|String|
|long_term_retention_backup_resource_id|String|
|auto_complete_restore|Bool|
|last_backup_name|String|
|location|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|