# Table: azure_network_watcher_flow_logs

This table shows data for Azure Network Watcher Flow Logs.

https://learn.microsoft.com/en-us/rest/api/network-watcher/flow-logs/list?tabs=HTTP#definitions

The primary key for this table is **id**.

## Relations

This table depends on [azure_network_watchers](azure_network_watchers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|id (PK)|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|