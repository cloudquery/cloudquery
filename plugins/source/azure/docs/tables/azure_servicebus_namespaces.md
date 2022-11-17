# Table: azure_servicebus_namespaces

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2#SBNamespace

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
|location|String|
|identity|JSON|
|alternate_name|String|
|disable_local_auth|Bool|
|encryption|JSON|
|minimum_tls_version|String|
|private_endpoint_connections|JSON|
|public_network_access|String|
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