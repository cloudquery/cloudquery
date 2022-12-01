# Table: azure_sql_servers

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql#Server

The primary key for this table is **id**.

## Relations

The following tables depend on azure_sql_servers:
  - [azure_sql_firewall_rules](azure_sql_firewall_rules.md)
  - [azure_sql_databases](azure_sql_databases.md)
  - [azure_sql_encryption_protectors](azure_sql_encryption_protectors.md)
  - [azure_sql_virtual_network_rules](azure_sql_virtual_network_rules.md)
  - [azure_sql_server_admins](azure_sql_server_admins.md)
  - [azure_sql_server_blob_auditing_policies](azure_sql_server_blob_auditing_policies.md)
  - [azure_sql_server_dev_ops_auditing_settings](azure_sql_server_dev_ops_auditing_settings.md)
  - [azure_sql_server_vulnerability_assessments](azure_sql_server_vulnerability_assessments.md)
  - [azure_sql_server_security_alert_policies](azure_sql_server_security_alert_policies.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|identity|JSON|
|kind|String|
|administrator_login|String|
|administrator_login_password|String|
|version|String|
|state|String|
|fully_qualified_domain_name|String|
|private_endpoint_connections|JSON|
|minimal_tls_version|String|
|public_network_access|String|
|location|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|