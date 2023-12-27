# Table: azure_network_virtual_networks

This table shows data for Azure Network Virtual Networks.

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/virtual-networks/list-all?tabs=HTTP#virtualnetwork

The primary key for this table is **id**.

## Relations

The following tables depend on azure_network_virtual_networks:
  - [azure_network_virtual_network_subnets](azure_network_virtual_network_subnets.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|extended_location|`json`|
|id (PK)|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|