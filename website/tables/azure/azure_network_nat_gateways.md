# Table: azure_network_nat_gateways

This table shows data for Azure Network Network Address Translation (NAT) Gateways.

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/nat-gateways/list?tabs=HTTP#natgateway

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|id (PK)|`utf8`|
|location|`utf8`|
|properties|`json`|
|sku|`json`|
|tags|`json`|
|zones|`list<item: utf8, nullable>`|
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|