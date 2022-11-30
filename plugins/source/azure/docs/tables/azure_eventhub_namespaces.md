# Table: azure_eventhub_namespaces

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub#EHNamespace

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
|sku|JSON|
|identity|JSON|
|provisioning_state|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|service_bus_endpoint|String|
|cluster_arm_id|String|
|metric_id|String|
|is_auto_inflate_enabled|Bool|
|maximum_throughput_units|Int|
|kafka_enabled|Bool|
|zone_redundant|Bool|
|encryption|JSON|
|location|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|