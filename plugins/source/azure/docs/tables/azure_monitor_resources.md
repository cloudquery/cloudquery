# Table: azure_monitor_resources

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources#GenericResourceExpanded

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
|id (PK)|String|