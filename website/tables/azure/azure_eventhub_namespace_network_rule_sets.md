# Table: azure_eventhub_namespace_network_rule_sets

This table shows data for Azure Event Hub Namespace Network Rule Sets.

https://learn.microsoft.com/en-us/rest/api/eventhub/stable/network-rule-sets/list-network-rule-set?tabs=HTTP#networkruleset

The primary key for this table is **id**.

## Relations

This table depends on [azure_eventhub_namespaces](azure_eventhub_namespaces).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|location|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|