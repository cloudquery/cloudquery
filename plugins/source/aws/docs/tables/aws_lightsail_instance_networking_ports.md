
# Table: aws_lightsail_instance_networking_ports
Describes information about ports for an Amazon Lightsail instance
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_instances table (FK)|
|access_direction|text|The access direction (inbound or outbound)|
|access_from|text|The location from which access is allowed|
|access_type|text|The type of access (Public or Private)|
|cidr_list_aliases|text[]|An alias that defines access for a preconfigured range of IP addresses|
|cidrs|text[]|The IPv4 address, or range of IPv4 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol|
|common_name|text|The common name of the port information|
|from_port|bigint|The first port in a range of open ports on an instance|
|ipv6_cidrs|text[]|The IPv6 address, or range of IPv6 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol|
|protocol|text|The IP protocol name|
|to_port|bigint|The last port in a range of open ports on an instance|
