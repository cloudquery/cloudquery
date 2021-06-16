
# Table: aws_ec2_route_table_routes
Describes a route in a route table.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|route_table_id|uuid|Unique ID of aws_ec2_route_tables table (FK)|
|carrier_gateway_id|text|The ID of the carrier gateway.|
|destination_cidr_block|text|The IPv4 CIDR block used for the destination match.|
|destination_ipv6_cidr_block|text|The IPv6 CIDR block used for the destination match.|
|destination_prefix_list_id|text|The prefix of the AWS service.|
|egress_only_internet_gateway_id|text|The ID of the egress-only internet gateway.|
|gateway_id|text|The ID of a gateway attached to your VPC.|
|instance_id|text|The ID of a NAT instance in your VPC.|
|instance_owner_id|text|The AWS account ID of the owner of the instance.|
|local_gateway_id|text|The ID of the local gateway.|
|nat_gateway_id|text|The ID of a NAT gateway.|
|network_interface_id|text|The ID of the network interface.|
|origin|text|Describes how the route was created.|
|state|text|The state of the route.|
|transit_gateway_id|text|The ID of a transit gateway.|
|vpc_peering_connection_id|text|The ID of a VPC peering connection.|
