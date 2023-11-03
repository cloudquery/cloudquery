# Table: azure_network_bgp_service_communities

This table shows data for Azure Network Border Gateway Protocol (BGP) Service Communities.

https://learn.microsoft.com/en-us/rest/api/expressroute/bgp-service-communities/list?tabs=HTTP#bgpservicecommunity

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id (PK)|`utf8`|
|id (PK)|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|name|`utf8`|
|type|`utf8`|