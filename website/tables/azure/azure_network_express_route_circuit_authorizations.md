# Table: azure_network_express_route_circuit_authorizations

This table shows data for Azure Network Express Route Circuit Authorizations.

https://learn.microsoft.com/en-us/rest/api/expressroute/express-route-circuit-authorizations/list?tabs=HTTP#expressroutecircuitauthorization

The primary key for this table is **id**.

## Relations

This table depends on [azure_network_express_route_circuits](azure_network_express_route_circuits).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|properties|`json`|
|etag|`utf8`|
|type|`utf8`|