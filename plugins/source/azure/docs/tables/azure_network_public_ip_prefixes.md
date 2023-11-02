# Table: azure_network_public_ip_prefixes

This table shows data for Azure Network Public IP Prefixes.

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/public-ip-prefixes/list?tabs=HTTP#publicipprefix

The primary key for this table is **id**.

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
|sku|`json`|
|tags|`json`|
|zones|`list<item: utf8, nullable>`|
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|