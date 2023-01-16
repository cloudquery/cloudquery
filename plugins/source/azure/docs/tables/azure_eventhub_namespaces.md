# Table: azure_eventhub_namespaces

The primary key for this table is **id**.

## Relations

The following tables depend on azure_eventhub_namespaces:
  - [azure_eventhub_network_rule_sets](azure_eventhub_network_rule_sets.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|id (PK)|String|
|identity|JSON|
|location|String|
|properties|JSON|
|sku|JSON|
|tags|JSON|
|name|String|
|system_data|JSON|
|type|String|