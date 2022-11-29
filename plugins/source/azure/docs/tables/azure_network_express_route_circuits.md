# Table: azure_network_express_route_circuits

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network#ExpressRouteCircuit

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|sku|JSON|
|allow_classic_operations|Bool|
|circuit_provisioning_state|String|
|service_provider_provisioning_state|String|
|authorizations|JSON|
|peerings|JSON|
|service_key|String|
|service_provider_notes|String|
|service_provider_properties|JSON|
|express_route_port|JSON|
|bandwidth_in_gbps|Float|
|stag|Int|
|provisioning_state|String|
|gateway_manager_etag|String|
|global_reach_enabled|Bool|
|etag|String|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|