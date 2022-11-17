# Table: azure_eventhub_network_rule_sets

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub#NetworkRuleSet

The primary key for this table is **id**.

## Relations
This table depends on [azure_eventhub_namespaces](azure_eventhub_namespaces.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|default_action|String|
|ip_rules|JSON|
|public_network_access|String|
|trusted_service_access_enabled|Bool|
|virtual_network_rules|JSON|
|id (PK)|String|
|location|String|
|name|String|
|system_data|JSON|
|type|String|
|namespace_id|String|