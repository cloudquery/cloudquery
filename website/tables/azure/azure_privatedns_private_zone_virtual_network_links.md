# Table: azure_privatedns_private_zone_virtual_network_links

This table shows data for Azure Privatedns Private Zone Virtual Network Links.

https://learn.microsoft.com/en-us/rest/api/dns/privatedns/virtual-network-links/list?tabs=HTTP#virtualnetworklink

The primary key for this table is **id**.

## Relations

This table depends on [azure_privatedns_private_zones](azure_privatedns_private_zones).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|etag|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|