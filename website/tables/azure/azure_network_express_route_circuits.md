# Table: azure_network_express_route_circuits

This table shows data for Azure Network Express Route Circuits.

https://learn.microsoft.com/en-us/rest/api/expressroute/express-route-circuits/list?tabs=HTTP#expressroutecircuit

The primary key for this table is **id**.

## Relations

The following tables depend on azure_network_express_route_circuits:
  - [azure_network_express_route_circuit_authorizations](azure_network_express_route_circuit_authorizations)
  - [azure_network_express_route_circuit_peerings](azure_network_express_route_circuit_peerings)

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
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|