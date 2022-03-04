
# Table: azure_eventhub_namespaces
Azure EventHub namespace
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|sku_name|text|Name of this SKU|
|sku_tier|text|The billing tier of this particular SKU|
|sku_capacity|integer|The Event Hubs throughput units, value should be 0 to 20 throughput units.|
|identity_principal_id|text|ObjectId from the KeyVault|
|identity_tenant_id|text|TenantId from the KeyVault|
|identity_type|text|Enumerates the possible value Identity type, which currently supports only 'SystemAssigned'|
|provisioning_state|text|Provisioning state of the Namespace.|
|created_at_time|timestamp without time zone||
|updated_at_time|timestamp without time zone||
|service_bus_endpoint|text|Endpoint you can use to perform Service Bus operations.|
|cluster_arm_id|text|Cluster ARM ID of the Namespace.|
|metric_id|text|Identifier for Azure Insights metrics.|
|is_auto_inflate_enabled|boolean|Value that indicates whether AutoInflate is enabled for eventhub namespace.|
|maximum_throughput_units|integer|Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20 throughput units|
|kafka_enabled|boolean|Value that indicates whether Kafka is enabled for eventhub namespace.|
|zone_redundant|boolean|Enabling this property creates a Standard Event Hubs Namespace in regions supported availability zones.|
|encryption_key_source|text|Enumerates the possible value of keySource for Encryption|
|location|text|Resource location.|
|tags|jsonb|Resource tags.|
|id|text|Resource ID.|
|name|text|Resource name.|
|type|text|Resource type.|
|network_rule_set|jsonb|Network rule set for a namespace.|
