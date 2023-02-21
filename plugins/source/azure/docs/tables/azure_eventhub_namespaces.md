# Table: azure_eventhub_namespaces

https://learn.microsoft.com/en-us/rest/api/eventhub/stable/namespaces/list?tabs=HTTP#ehnamespace

The primary key for this table is **id**.

## Relations

The following tables depend on azure_eventhub_namespaces:
  - [azure_eventhub_namespace_network_rule_sets](azure_eventhub_namespace_network_rule_sets.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|identity|JSON|
|location|String|
|properties|JSON|
|sku|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|