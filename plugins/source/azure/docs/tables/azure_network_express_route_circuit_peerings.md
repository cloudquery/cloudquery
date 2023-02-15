# Table: azure_network_express_route_circuit_peerings

https://learn.microsoft.com/en-us/rest/api/expressroute/express-route-circuit-peerings/list?tabs=HTTP#expressroutecircuitpeering

The composite primary key for this table is (**express_route_circuit_name**, **id**).

## Relations

This table depends on [azure_network_express_route_circuits](azure_network_express_route_circuits.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|express_route_circuit_name (PK)|String|
|id (PK)|String|
|name|String|
|properties|JSON|
|etag|String|
|type|String|