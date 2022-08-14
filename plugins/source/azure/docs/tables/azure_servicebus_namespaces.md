
# Table: azure_servicebus_namespaces
SBNamespace description of a namespace resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|sku_name|text|Name of this SKU|
|sku_tier|text|The billing tier of this particular SKU|
|sku_capacity|integer|The specified messaging units for the tier|
|identity_principal_id|text|ObjectId from the KeyVault|
|identity_tenant_id|text|TenantId from the KeyVault|
|identity_type|text|Type of managed service identity|
|user_assigned_identities|jsonb|Properties for User Assigned Identities|
|system_data|jsonb|The system meta data relating to this resource|
|provisioning_state|text|Provisioning state of the namespace|
|status|text|Status of the namespace|
|created_at_time|timestamp without time zone||
|updated_at_time|timestamp without time zone||
|service_bus_endpoint|text|Endpoint you can use to perform Service Bus operations|
|metric_id|text|Identifier for Azure Insights metrics|
|zone_redundant|boolean|Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones|
|key_vault_properties|jsonb|Properties of KeyVault|
|key_source|text|Enumerates the possible value of keySource for Encryption|
|require_infrastructure_encryption|boolean|Enable Infrastructure Encryption (Double Encryption)|
|disable_local_auth|boolean|This property disables SAS authentication for the Service Bus namespace|
|location|text|The Geo-location where the resource lives|
|tags|jsonb|Resource tags|
|id|text|Resource Id|
|name|text|Resource name|
|type|text|Resource type|
