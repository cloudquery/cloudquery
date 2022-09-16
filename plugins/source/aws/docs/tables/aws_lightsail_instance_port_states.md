
# Table: aws_lightsail_instance_port_states
Describes open ports on an instance, the IP addresses allowed to connect to the instance through the ports, and the protocol
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_instances table (FK)|
|cidr_list_aliases|text[]|An alias that defines access for a preconfigured range of IP addresses|
|cidrs|text[]|The IPv4 address, or range of IPv4 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol|
|from_port|bigint|The first port in a range of open ports on an instance|
|ipv6_cidrs|text[]|The IPv6 address, or range of IPv6 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol|
|protocol|text|The IP protocol name|
|state|text|Specifies whether the instance port is open or closed|
|to_port|bigint|The last port in a range of open ports on an instance|
