
# Table: azure_search_services
Service describes an Azure Cognitive Search service and its current state.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|replica_count|integer|The number of replicas in the search service|
|partition_count|integer|The number of partitions in the search service; if specified, it can be 1, 2, 3, 4, 6, or 12|
|hosting_mode|text|Applicable only for the standard3 SKU|
|public_network_access|text|This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates|
|status|text|The status of the search service|
|status_details|text|The details of the search service status.|
|provisioning_state|text|The state of the last provisioning operation performed on the search service|
|network_rule_set_ip_rules|inet[]|A list of IP restriction rules that defines the inbound network(s) with allowing access to the search service endpoint|
|sku_name|text|The SKU of the search service|
|identity_principal_id|text|The principal ID of resource identity.|
|identity_tenant_id|text|The tenant ID of resource.|
|identity_type|text|The identity type|
|tags|jsonb|Resource tags.|
|location|text|The geo-location where the resource lives|
|id|text|Fully qualified resource ID for the resource|
|name|text|The name of the resource|
|type|text|The type of the resource|
