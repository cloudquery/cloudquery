# Table: azure_sql_servers

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql#Server

The primary key for this table is **id**.

## Relations

The following tables depend on azure_sql_servers:
  - [azure_sql_databases](azure_sql_databases.md)
  - [azure_sql_encryption_protectors](azure_sql_encryption_protectors.md)
  - [azure_sql_firewall_rules](azure_sql_firewall_rules.md)
  - [azure_sql_server_administrators](azure_sql_server_administrators.md)
  - [azure_sql_server_blob_auditing_policies](azure_sql_server_blob_auditing_policies.md)
  - [azure_sql_server_dev_ops_auditing_settings](azure_sql_server_dev_ops_auditing_settings.md)
  - [azure_sql_server_security_alert_policies](azure_sql_server_security_alert_policies.md)
  - [azure_sql_server_vulnerability_assessments](azure_sql_server_vulnerability_assessments.md)
  - [azure_sql_virtual_network_rules](azure_sql_virtual_network_rules.md)

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
|federated_client_id|String|
|key_id|String|
|minimal_tls_version|String|
|primary_user_assigned_identity_id|String|
|public_network_access|String|
|restrict_outbound_network_access|String|
|version|String|
|fully_qualified_domain_name|String|
|private_endpoint_connections|JSON|
|state|String|
|workspace_feature|String|
|tags|JSON|
|id (PK)|String|
|kind|String|
|name|String|
|type|String|