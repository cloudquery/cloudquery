# Table: azure_monitor_resources

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources#GenericResourceExpanded

The primary key for this table is **id**.

## Relations

The following tables depend on azure_monitor_resources:
  - [azure_monitor_diagnostic_settings](azure_monitor_diagnostic_settings.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|extended_location|JSON|
|identity|JSON|
|kind|String|
|location|String|
|managed_by|String|
|plan|JSON|
|sku|JSON|
|tags|JSON|
|changed_time|Timestamp|
|created_time|Timestamp|
|id (PK)|String|
|name|String|
|provisioning_state|String|
|type|String|