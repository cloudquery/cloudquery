# Table: azure_network_virtual_network_subnets

This table shows data for Azure Network Virtual Network Subnets.

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/subnets/list?tabs=HTTP#subnet

The primary key for this table is **id**.

## Relations

This table depends on [azure_network_virtual_networks](azure_network_virtual_networks).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|properties|`json`|
|type|`utf8`|
|etag|`utf8`|