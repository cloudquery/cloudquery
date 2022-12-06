# Table: azure_network_flow_logs

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network#FlowLog

The primary key for this table is **id**.

## Relations
This table depends on [azure_network_watchers](azure_network_watchers.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|network_watcher_id|String|
|target_resource_id|String|
|target_resource_guid|String|
|storage_id|String|
|enabled|Bool|
|retention_policy|JSON|
|format|JSON|
|flow_analytics_configuration|JSON|
|provisioning_state|String|
|etag|String|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|