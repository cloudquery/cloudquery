# Table: azure_network_express_route_circuit_authorizations

https://learn.microsoft.com/en-us/rest/api/expressroute/express-route-circuit-authorizations/list?tabs=HTTP#expressroutecircuitauthorization

The primary key for this table is **id**.

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
|id (PK)|String|
|name|String|
|properties|JSON|
|etag|String|
|type|String|