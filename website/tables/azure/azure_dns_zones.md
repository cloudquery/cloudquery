# Table: azure_dns_zones

This table shows data for Azure DNS Zones.

https://learn.microsoft.com/en-us/rest/api/dns/zones/list?tabs=HTTP#zone

The primary key for this table is **id**.

## Relations

The following tables depend on azure_dns_zones:
  - [azure_dns_record_sets](azure_dns_record_sets)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|etag|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|