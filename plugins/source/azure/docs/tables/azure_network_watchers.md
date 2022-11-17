# Table: azure_network_watchers

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2#Watcher

The primary key for this table is **id**.

## Relations

The following tables depend on azure_network_watchers:
  - [azure_network_flow_logs](azure_network_flow_logs.md)

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
|provisioning_state|String|
|tags|JSON|
|etag|String|
|name|String|
|type|String|