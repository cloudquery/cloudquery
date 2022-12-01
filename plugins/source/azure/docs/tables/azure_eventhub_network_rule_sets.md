# Table: azure_eventhub_network_rule_sets

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub#NetworkRuleSet

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
|eventhub_namespace_id|String|
|trusted_service_access_enabled|Bool|
|default_action|String|
|virtual_network_rules|JSON|
|ip_rules|JSON|
|id (PK)|String|
|name|String|
|type|String|