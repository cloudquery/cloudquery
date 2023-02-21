# Table: azure_network_virtual_network_subnets

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/subnets/list?tabs=HTTP#subnet

The primary key for this table is **id**.

## Relations

This table depends on [azure_network_virtual_networks](azure_network_virtual_networks.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|id (PK)|String|
|name|String|
|properties|JSON|
|type|String|
|etag|String|