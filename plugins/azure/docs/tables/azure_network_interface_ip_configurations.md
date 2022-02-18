
# Table: azure_network_interface_ip_configurations
NetworkInterface IP Configurations. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|interface_cq_id|uuid|Unique CloudQuery ID of azure_network_interface table (FK)|
|id|text|Resource Id|
|name|text|Resource name|
|etag|text|A unique read-only string that changes whenever the resource is updated.|
|type|text|Resource type|
|primary|boolean|Whether this is a primary network interface on a virtual machine.|
|application_gateway_backend_address_pools|jsonb|The reference to ApplicationGatewayBackendAddressPool resource.|
|application_security_groups|jsonb|Application security groups in which the IP configuration is included.|
|load_balancer_backend_address_pools|jsonb|The reference to LoadBalancerBackendAddressPool resource.|
|load_balancer_inbound_nat_rules|jsonb|A list of references of LoadBalancerInboundNatRules.|
|private_ip_address|text|Private IP address of the IP configuration.|
|private_ip_address_version|text|Whether the specific IP configuration is IPv4 or IPv6. Default is IPv4. Possible values include: 'IPVersionIPv4', 'IPVersionIPv6|
|private_ip_allocation_method|text|Private IP address allocation method.|
|private_link_connection_properties|text|PrivateLinkConnection properties for the network interface.|
|provisioning_state|text|The provisioning state of the network interface IP configuration.|
|public_ip_address|text|Public IP address bound to the IP configuration.|
|subnet_id|text|subnet ID of network interface ip configuration|
|virtual_network_taps|jsonb|The reference to Virtual Network Taps.|
