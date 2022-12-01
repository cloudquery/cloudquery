# Table: azure_mariadb_servers

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2020-01-01/mariadb#Server

The primary key for this table is **id**.

## Relations

The following tables depend on azure_mariadb_servers:
  - [azure_mariadb_configurations](azure_mariadb_configurations.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|sku|JSON|
|administrator_login|String|
|version|String|
|ssl_enforcement|String|
|user_visible_state|String|
|fully_qualified_domain_name|String|
|earliest_restore_date|Timestamp|
|storage_profile|JSON|
|replication_role|String|
|master_server_id|String|
|replica_capacity|Int|
|public_network_access|String|
|private_endpoint_connections|JSON|
|tags|JSON|
|location|String|
|id (PK)|String|
|name|String|
|type|String|