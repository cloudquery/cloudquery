
# Table: azure_network_virtual_networks
Azure virtual network
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|extended_location_name|text|The name of the extended location.|
|extended_location_type|text|The type of the extended location.|
|address_space_address_prefixes|text[]|A list of address blocks reserved for this virtual network in CIDR notation.|
|dhcp_options_dns_servers|inet[]|The list of DNS servers IP addresses.|
|resource_guid|text|The resourceGuid property of the Virtual Network resource.|
|provisioning_state|text|The provisioning state of the virtual network resource|
|enable_ddos_protection|boolean|Indicates if DDoS protection is enabled for all the protected resources in the virtual network|
|enable_vm_protection|boolean|Indicates if VM protection is enabled for all the subnets in the virtual network.|
|ddos_protection_plan_id|text|Resource ID.|
|bgp_communities_virtual_network_community|text|The BGP community associated with the virtual network.|
|bgp_communities_regional_community|text|The BGP community associated with the region of the virtual network.|
|ip_allocations|text[]|Array of IpAllocation which reference this VNET.|
|etag|text|A unique read-only string that changes whenever the resource is updated.|
|id|text|Resource ID.|
|name|text|Resource name.|
|type|text|Resource type.|
|location|text|Resource location.|
|tags|jsonb|Resource tags.|
