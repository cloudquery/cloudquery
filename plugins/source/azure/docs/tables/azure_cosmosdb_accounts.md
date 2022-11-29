# Table: azure_cosmosdb_accounts

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb#DatabaseAccountGetResults

The primary key for this table is **id**.

## Relations

The following tables depend on azure_cosmosdb_accounts:
  - [azure_cosmosdb_mongo_db_databases](azure_cosmosdb_mongo_db_databases.md)
  - [azure_cosmosdb_sql_databases](azure_cosmosdb_sql_databases.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|kind|String|
|provisioning_state|String|
|document_endpoint|String|
|database_account_offer_type|String|
|ip_rules|JSON|
|is_virtual_network_filter_enabled|Bool|
|enable_automatic_failover|Bool|
|consistency_policy|JSON|
|capabilities|JSON|
|write_locations|JSON|
|read_locations|JSON|
|locations|JSON|
|failover_policies|JSON|
|virtual_network_rules|JSON|
|private_endpoint_connections|JSON|
|enable_multiple_write_locations|Bool|
|enable_cassandra_connector|Bool|
|connector_offer|String|
|disable_key_based_metadata_write_access|Bool|
|key_vault_key_uri|String|
|public_network_access|String|
|enable_free_tier|Bool|
|api_properties|JSON|
|enable_analytical_storage|Bool|
|cors|JSON|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|