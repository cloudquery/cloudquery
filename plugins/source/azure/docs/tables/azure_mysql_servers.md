# Table: azure_mysql_servers

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql#Server

The primary key for this table is **id**.

## Relations

The following tables depend on azure_mysql_servers:
  - [azure_mysql_configurations](azure_mysql_configurations.md)

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
|earliest_restore_date|Timestamp|
|fully_qualified_domain_name|String|
|infrastructure_encryption|String|
|master_server_id|String|
|minimal_tls_version|String|
|public_network_access|String|
|replica_capacity|Int|
|replication_role|String|
|ssl_enforcement|String|
|storage_profile|JSON|
|user_visible_state|String|
|version|String|
|byok_enforcement|String|
|private_endpoint_connections|JSON|
|sku|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|