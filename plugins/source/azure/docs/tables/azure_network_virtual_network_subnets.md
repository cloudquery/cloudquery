# Table: azure_network_virtual_network_subnets

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/subnets/list?tabs=HTTP#subnet

The composite primary key for this table is (**virtual_network_name**, **id**).

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
|virtual_network_name (PK)|String|
|id (PK)|String|
|name|String|
|properties|JSON|
|type|String|
|etag|String|