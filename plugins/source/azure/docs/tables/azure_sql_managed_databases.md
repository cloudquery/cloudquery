# Table: azure_sql_managed_databases

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql#ManagedDatabase

The primary key for this table is **id**.

## Relations
This table depends on [azure_sql_managed_instances](azure_sql_managed_instances.md).

The following tables depend on azure_sql_managed_databases:
  - [azure_sql_managed_database_vulnerability_assessment_scans](azure_sql_managed_database_vulnerability_assessment_scans.md)
  - [azure_sql_managed_database_vulnerability_assessments](azure_sql_managed_database_vulnerability_assessments.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|auto_complete_restore|Bool|
|catalog_collation|String|
|collation|String|
|create_mode|String|
|last_backup_name|String|
|long_term_retention_backup_resource_id|String|
|recoverable_database_id|String|
|restorable_dropped_database_id|String|
|restore_point_in_time|Timestamp|
|source_database_id|String|
|storage_container_sas_token|String|
|storage_container_uri|String|
|creation_date|Timestamp|
|default_secondary_location|String|
|earliest_restore_point|Timestamp|
|failover_group_id|String|
|status|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|
|managed_instance_id|String|