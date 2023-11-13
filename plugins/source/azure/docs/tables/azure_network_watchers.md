# Table: azure_network_watchers

This table shows data for Azure Network Watchers.

https://learn.microsoft.com/en-us/rest/api/network-watcher/network-watchers/list-all?tabs=HTTP#networkwatcher

The primary key for this table is **id**.

## Relations

The following tables depend on azure_network_watchers:
  - [azure_network_watcher_flow_logs](azure_network_watcher_flow_logs.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
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