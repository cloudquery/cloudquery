
# Table: azure_network_express_route_circuits
Azure Network Express Route Circuits
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id.|
|id|text|Resource ID.|
|allow_classic_operations|boolean|Allow classic operations.|
|bandwidth_in_gbps|float|The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.|
|circuit_provisioning_state|text|The CircuitProvisioningState state of the resource.|
|etag|text|A unique read-only string that changes whenever the resource is updated.|
|express_route_port_id|text|The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.|
|gateway_manager_etag|text|The GatewayManager Etag.|
|global_reach_enabled|boolean|Flag denoting global reach status.|
|location|text|Resource location.|
|name|text|Resource name.|
|provisioning_state|text|The provisioning state of the express route circuit resource.|
|service_key|text|The ServiceKey.|
|service_provider_notes|text|The ServiceProviderNotes.|
|service_provider_properties_bandwidth_in_mbps|integer|The BandwidthInMbps.|
|service_provider_properties_peering_location|text|The peering location.|
|service_provider_properties_service_provider_name|text|The serviceProviderName.|
|service_provider_provisioning_state|text|The ServiceProviderProvisioningState state of the resource.|
|sku_family|text|The family of the SKU.|
|sku_name|text|The name of the SKU.|
|sku_tier|text|The tier of the SKU.|
|stag|integer|The identifier of the circuit traffic. Outer tag for QinQ encapsulation.|
|tags|jsonb|Resource tags.|
|type|text|Resource type.|
