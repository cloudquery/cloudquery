# Table: azure_network_flow_logs

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2#FlowLog

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
|id (PK)|String|
|location|String|
|storage_id|String|
|target_resource_id|String|
|enabled|Bool|
|flow_analytics_configuration|JSON|
|format|JSON|
|retention_policy|JSON|
|provisioning_state|String|
|target_resource_guid|String|
|tags|JSON|
|etag|String|
|name|String|
|type|String|
|watcher_id|String|