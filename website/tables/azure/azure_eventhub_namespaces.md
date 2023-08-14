# Table: azure_eventhub_namespaces

This table shows data for Azure Event Hub Namespaces.

https://learn.microsoft.com/en-us/rest/api/eventhub/stable/namespaces/list?tabs=HTTP#ehnamespace

The primary key for this table is **id**.

## Relations

The following tables depend on azure_eventhub_namespaces:
  - [azure_eventhub_namespace_network_rule_sets](azure_eventhub_namespace_network_rule_sets)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|identity|`json`|
|location|`utf8`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|