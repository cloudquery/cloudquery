# Table: azure_mariadb_servers

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb#Server

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
|location|String|
|administrator_login|String|
|earliest_restore_date|Timestamp|
|fully_qualified_domain_name|String|
|master_server_id|String|
|minimal_tls_version|String|
|public_network_access|String|
|replica_capacity|Int|
|replication_role|String|
|ssl_enforcement|String|
|storage_profile|JSON|
|user_visible_state|String|
|version|String|
|private_endpoint_connections|JSON|
|sku|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|