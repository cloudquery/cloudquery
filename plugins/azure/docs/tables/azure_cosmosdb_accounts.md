
# Table: azure_cosmosdb_accounts
Azure Cosmos DB database account.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|provisioning_state|text||
|document_endpoint|text|The connection endpoint for the Cosmos DB database account.|
|database_account_offer_type|text|The offer type for the Cosmos DB database account|
|ip_rules|text[]|List of IpRules.|
|capabilities|text[]|Capability cosmos DB capability object|
|is_virtual_network_filter_enabled|boolean|Flag to indicate whether to enable/disable Virtual Network ACL rules.|
|enable_automatic_failover|boolean|Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage|
|consistency_policy_default_consistency_level|text|The default consistency level and configuration settings of the Cosmos DB account|
|consistency_policy_max_staleness_prefix|bigint|When used with the Bounded Staleness consistency level, this value represents the number of stale requests tolerated|
|consistency_policy_max_interval_in_seconds|integer|When used with the Bounded Staleness consistency level, this value represents the time amount of staleness (in seconds) tolerated|
|virtual_network_rules|jsonb|List of Virtual Network ACL rules configured for the Cosmos DB account.|
|enable_multiple_write_locations|boolean|Enables the account to write in multiple locations|
|enable_cassandra_connector|boolean|Enables the cassandra connector on the Cosmos DB C* account|
|connector_offer|text|The cassandra connector offer type for the Cosmos DB database C* account|
|disable_key_based_metadata_write_access|boolean|Disable write operations on metadata resources (databases, containers, throughput) via account keys|
|key_vault_key_uri|text|The URI of the key vault|
|public_network_access|text|Whether requests from Public Network are allowed|
|enable_free_tier|boolean|Flag to indicate whether Free Tier is enabled.|
|api_properties_server_version|text|Describes the ServerVersion of an a MongoDB account|
|enable_analytical_storage|boolean|Flag to indicate whether to enable storage analytics.|
|id|text|The unique resource identifier of the ARM resource.|
|name|text|The name of the ARM resource.|
|type|text|The type of Azure resource.|
|location|text|The location of the resource group to which the resource belongs.|
|tags|jsonb|Resource tags.|
