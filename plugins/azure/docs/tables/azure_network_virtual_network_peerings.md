
# Table: azure_network_virtual_network_peerings
Azure virtual network peering
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_network_cq_id|uuid|Unique CloudQuery ID of azure_network_virtual_networks table (FK)|
|allow_virtual_network_access|boolean|Whether the VMs in the local virtual network space would be able to access the VMs in remote virtual network space|
|allow_forwarded_traffic|boolean|Whether the forwarded traffic from the VMs in the local virtual network will be allowed/disallowed in remote virtual network|
|allow_gateway_transit|boolean|If gateway links can be used in remote virtual networking to link to this virtual network|
|use_remote_gateways|boolean|If remote gateways can be used on this virtual network If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit Only one peering can have this flag set to true This flag cannot be set if virtual network already has a gateway|
|remote_virtual_network_id|text|Resource ID|
|remote_address_space_address_prefixes|text[]|A list of address blocks reserved for this virtual network in CIDR notation|
|remote_bgp_communities_virtual_network_community|text|The BGP community associated with the virtual network|
|remote_bgp_communities_regional_community|text|The BGP community associated with the region of the virtual network|
|peering_state|text|The status of the virtual network peering Possible values include: 'VirtualNetworkPeeringStateInitiated', 'VirtualNetworkPeeringStateConnected', 'VirtualNetworkPeeringStateDisconnected'|
|provisioning_state|text|The provisioning state of the virtual network peering resource Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'|
|name|text|The name of the resource that is unique within a resource group This name can be used to access the resource|
|etag|text|A unique read-only string that changes whenever the resource is updated|
|id|text|Resource ID|
