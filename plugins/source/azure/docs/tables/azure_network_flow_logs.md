# Table: azure_network_flow_logs


The primary key for this table is **id**.

## Relations
This table depends on [`azure_network_watchers`](azure_network_watchers.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|network_watcher_id|UUID|
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|