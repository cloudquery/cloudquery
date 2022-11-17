# Table: azure_sql_managed_instances

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql#ManagedInstance

The primary key for this table is **id**.

## Relations

The following tables depend on azure_sql_managed_instances:
  - [azure_sql_managed_databases](azure_sql_managed_databases.md)
  - [azure_sql_managed_instance_encryption_protectors](azure_sql_managed_instance_encryption_protectors.md)
  - [azure_sql_managed_instance_vulnerability_assessments](azure_sql_managed_instance_vulnerability_assessments.md)

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
|administrator_login|String|
|administrator_login_password|String|
|administrators|JSON|
|collation|String|
|dns_zone_partner|String|
|instance_pool_id|String|
|key_id|String|
|license_type|String|
|maintenance_configuration_id|String|
|managed_instance_create_mode|String|
|minimal_tls_version|String|
|primary_user_assigned_identity_id|String|
|proxy_override|String|
|public_data_endpoint_enabled|Bool|
|requested_backup_storage_redundancy|String|
|restore_point_in_time|Timestamp|
|service_principal|JSON|
|source_managed_instance_id|String|
|storage_size_in_gb|Int|
|subnet_id|String|
|timezone_id|String|
|v_cores|Int|
|zone_redundant|Bool|
|current_backup_storage_redundancy|String|
|dns_zone|String|
|fully_qualified_domain_name|String|
|private_endpoint_connections|JSON|
|provisioning_state|String|
|state|String|
|sku|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|