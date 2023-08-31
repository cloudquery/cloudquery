# Table: azure_privatedns_private_zones

This table shows data for Azure Privatedns Private Zones.

https://learn.microsoft.com/en-us/rest/api/dns/privatedns/private-zones/list?tabs=HTTP#privatezone

The primary key for this table is **id**.

## Relations

The following tables depend on azure_privatedns_private_zones:
  - [azure_privatedns_private_zone_record_sets](azure_privatedns_private_zone_record_sets)
  - [azure_privatedns_private_zone_virtual_network_links](azure_privatedns_private_zone_virtual_network_links)

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