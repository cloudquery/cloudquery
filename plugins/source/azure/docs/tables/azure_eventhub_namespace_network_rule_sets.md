# Table: azure_eventhub_namespace_network_rule_sets

https://learn.microsoft.com/en-us/rest/api/eventhub/stable/network-rule-sets/list-network-rule-set?tabs=HTTP#networkruleset

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
|properties|JSON|
|id (PK)|String|
|location|String|
|name|String|
|system_data|JSON|
|type|String|