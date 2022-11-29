# Table: azure_servicebus_namespaces

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus#SBNamespace

The primary key for this table is **id**.

## Relations

The following tables depend on azure_servicebus_namespaces:
  - [azure_servicebus_topics](azure_servicebus_topics.md)

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
|system_data|JSON|
|provisioning_state|String|
|status|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|service_bus_endpoint|String|
|metric_id|String|
|zone_redundant|Bool|
|encryption|JSON|
|private_endpoint_connections|JSON|
|disable_local_auth|Bool|
|location|String|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|