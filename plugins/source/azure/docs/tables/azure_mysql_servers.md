# Table: azure_mysql_servers



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
|identity|JSON|
|sku|JSON|
|administrator_login|String|
|version|String|
|ssl_enforcement|String|
|minimal_tls_version|String|
|byok_enforcement|String|
|infrastructure_encryption|String|
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