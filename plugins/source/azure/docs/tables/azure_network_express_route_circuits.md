# Table: azure_network_express_route_circuits

https://learn.microsoft.com/en-us/rest/api/expressroute/express-route-circuits/list?tabs=HTTP#expressroutecircuit

The primary key for this table is **id**.

## Relations

The following tables depend on azure_network_express_route_circuits:
  - [azure_network_express_route_circuit_authorizations](azure_network_express_route_circuit_authorizations.md)
  - [azure_network_express_route_circuit_peerings](azure_network_express_route_circuit_peerings.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|id (PK)|String|
|location|String|
|properties|JSON|
|sku|JSON|
|tags|JSON|
|etag|String|
|name|String|
|type|String|