# Table: azure_eventhub_network_rule_sets


The primary key for this table is **id**.

## Relations
This table depends on [`azure_eventhub_namespaces`](azure_eventhub_namespaces.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|eventhub_namespace_id|UUID|
|trusted_service_access_enabled|Bool|
|default_action|String|
|virtual_network_rules|JSON|
|ip_rules|JSON|
|id (PK)|String|
|name|String|
|type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|