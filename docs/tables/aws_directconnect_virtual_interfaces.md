
# Table: aws_directconnect_virtual_interfaces
Information about a virtual interface. A virtual interface (VLAN) transmits the traffic between the AWS Direct Connect location and the customer network
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|address_family|text|The address family for the BGP peer.|
|amazon_address|text|The IP address assigned to the Amazon interface.|
|amazon_side_asn|bigint|The autonomous system number (ASN) for the Amazon side of the connection.|
|asn|integer|The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration|
|auth_key|text|The authentication key for BGP configuration|
|aws_device_v2|text|The Direct Connect endpoint on which the virtual interface terminates.|
|connection_id|text|The ID of the connection.|
|customer_address|text|The IP address assigned to the customer interface.|
|customer_router_config|text|The customer router configuration.|
|direct_connect_gateway_id|text|The ID of the Direct Connect gateway.|
|jumbo_frame_capable|boolean|Indicates whether jumbo frames (9001 MTU) are supported.|
|location|text|The location of the connection.|
|mtu|integer|The maximum transmission unit (MTU), in bytes|
|owner_account|text|The ID of the AWS account that owns the virtual interface.|
|region|text|The AWS Region where the virtual interface is located.|
|route_filter_prefixes|text[]|The routes to be advertised to the AWS network in this Region|
|tags|jsonb|The tags associated with the virtual interface.|
|virtual_gateway_id|text|The ID of the virtual private gateway|
|id|text|The ID of the virtual interface.|
|virtual_interface_name|text|The name of the virtual interface assigned by the customer network|
|virtual_interface_state|text|The state of the virtual interface|
|virtual_interface_type|text|The type of virtual interface|
|vlan|integer|The ID of the VLAN.|
