
# Table: azure_network_express_route_ports
Azure Network Express Route Ports
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id.|
|id|text|Resource ID.|
|allocation_date|text|Date of the physical port allocation to be used in Letter of Authorization.|
|bandwidth_in_gbps|integer|Bandwidth of procured ports in Gbps.|
|circuits|text[]|Reference the ExpressRoute circuit(s) that are provisioned on this ExpressRoutePort resource.|
|encapsulation|text|Encapsulation method on physical ports.|
|etag|text|A unique read-only string that changes whenever the resource is updated.|
|ether_type|text|Ether type of the physical port.|
|identity_principal_id|text|The principal id of the system assigned identity.|
|identity_tenant_id|text|The tenant id of the system assigned identity.|
|identity_type|text|The type of identity used for the resource.|
|identity_user_assigned_identities|jsonb|The list of user identities associated with resource.|
|location|text|Resource location.|
|mtu|text|Maximum transmission unit of the physical port pair(s).|
|name|text|Resource name.|
|peering_location|text|The name of the peering location that the ExpressRoutePort is mapped to physically.|
|provisioned_bandwidth_in_gbps|float|Aggregate Gbps of associated circuit bandwidths.|
|provisioning_state|text|The provisioning state of the express route port resource.|
|resource_guid|text|The resource GUID property of the express route port resource.|
|tags|jsonb|Resource tags.|
|type|text|Resource type.|
