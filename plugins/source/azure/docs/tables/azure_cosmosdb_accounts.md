# Table: azure_cosmosdb_accounts

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2#DatabaseAccountGetResults

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
|identity|JSON|
|kind|String|
|location|String|
|api_properties|JSON|
|analytical_storage_configuration|JSON|
|capabilities|JSON|
|capacity|JSON|
|connector_offer|String|
|consistency_policy|JSON|
|cors|JSON|
|create_mode|String|
|default_identity|String|
|disable_key_based_metadata_write_access|Bool|
|disable_local_auth|Bool|
|enable_analytical_storage|Bool|
|enable_automatic_failover|Bool|
|enable_cassandra_connector|Bool|
|enable_free_tier|Bool|
|enable_multiple_write_locations|Bool|
|enable_partition_merge|Bool|
|ip_rules|JSON|
|is_virtual_network_filter_enabled|Bool|
|key_vault_key_uri|String|
|network_acl_bypass|String|
|network_acl_bypass_resource_ids|StringArray|
|public_network_access|String|
|restore_parameters|JSON|
|virtual_network_rules|JSON|
|database_account_offer_type|String|
|document_endpoint|String|
|failover_policies|JSON|
|instance_id|String|
|keys_metadata|JSON|
|locations|JSON|
|private_endpoint_connections|JSON|
|provisioning_state|String|
|read_locations|JSON|
|write_locations|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|