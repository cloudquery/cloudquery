# Table: azure_sql_managed_instances

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql#ManagedInstance

The primary key for this table is **id**.

## Relations

The following tables depend on azure_sql_managed_instances:
  - [azure_sql_managed_databases](azure_sql_managed_databases.md)
  - [azure_sql_managed_instance_vulnerability_assessments](azure_sql_managed_instance_vulnerability_assessments.md)
  - [azure_sql_managed_instance_encryption_protectors](azure_sql_managed_instance_encryption_protectors.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|identity|JSON|
|sku|JSON|
|provisioning_state|String|
|managed_instance_create_mode|String|
|fully_qualified_domain_name|String|
|administrator_login|String|
|administrator_login_password|String|
|subnet_id|String|
|state|String|
|license_type|String|
|v_cores|Int|
|storage_size_in_gb|Int|
|collation|String|
|dns_zone|String|
|dns_zone_partner|String|
|public_data_endpoint_enabled|Bool|
|source_managed_instance_id|String|
|restore_point_in_time|Timestamp|
|proxy_override|String|
|timezone_id|String|
|instance_pool_id|String|
|maintenance_configuration_id|String|
|private_endpoint_connections|JSON|
|minimal_tls_version|String|
|storage_account_type|String|
|zone_redundant|Bool|
|location|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|