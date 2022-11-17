# Table: azure_eventhub_namespaces

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub#EHNamespace

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
|identity|JSON|
|location|String|
|alternate_name|String|
|cluster_arm_id|String|
|disable_local_auth|Bool|
|encryption|JSON|
|is_auto_inflate_enabled|Bool|
|kafka_enabled|Bool|
|maximum_throughput_units|Int|
|private_endpoint_connections|JSON|
|zone_redundant|Bool|
|created_at|Timestamp|
|metric_id|String|
|provisioning_state|String|
|service_bus_endpoint|String|
|status|String|
|updated_at|Timestamp|
|sku|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|