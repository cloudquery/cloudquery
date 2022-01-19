
# Table: azure_network_virtual_network_subnets
Azure virtual network subnet
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_network_cq_id|uuid|Unique CloudQuery ID of azure_network_virtual_networks table (FK)|
|address_prefix|text|The address prefix for the subnet.|
|address_prefixes|text[]|List of address prefixes for the subnet.|
|security_group_properties_format_resource_guid|text|The resource GUID property of the network security group resource.|
|security_group_properties_format_provisioning_state|text|The provisioning state of the network security group resource|
|network_security_group_etag|text|A unique read-only string that changes whenever the resource is updated.|
|network_security_group_id|text|Resource ID.|
|network_security_group_name|text|Resource name.|
|network_security_group_type|text|Resource type.|
|network_security_group_location|text|Resource location.|
|network_security_group_tags|jsonb|Resource tags.|
|route_table_disable_bgp_route_propagation|boolean|Whether to disable the routes learned by BGP on that route table|
|route_table_provisioning_state|text|The provisioning state of the route table resource|
|route_table_resource_guid|text|The resource GUID property of the route table.|
|route_table_etag|text|A unique read-only string that changes whenever the resource is updated.|
|route_table_id|text|Resource ID.|
|route_table_name|text|Resource name.|
|route_table_type|text|Resource type.|
|route_table_location|text|Resource location.|
|route_table_tags|jsonb|Resource tags.|
|nat_gateway_id|text|Resource ID.|
|purpose|text|A read-only string identifying the intention of use for this subnet based on delegations and other user-defined properties.|
|provisioning_state|text|The provisioning state of the subnet resource|
|private_endpoint_network_policies|text|Enable or Disable apply network policies on private end point in the subnet.|
|private_link_service_network_policies|text|Enable or Disable apply network policies on private link service in the subnet.|
|name|text|The name of the resource that is unique within a resource group|
|etag|text|A unique read-only string that changes whenever the resource is updated.|
|id|text|Resource ID.|
