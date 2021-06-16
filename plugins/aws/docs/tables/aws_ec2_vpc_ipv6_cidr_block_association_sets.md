
# Table: aws_ec2_vpc_ipv6_cidr_block_association_sets
Describes an IPv6 CIDR block associated with a VPC.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vpc_id|uuid|Unique ID of aws_ec2_vpcs table (FK)|
|association_id|text|The association ID for the IPv6 CIDR block.|
|ipv6_cidr_block|text|The IPv6 CIDR block.|
|ipv6_cidr_block_state|text|The state of the CIDR block.|
|ipv6_cidr_block_state_status_message|text|A message about the status of the CIDR block, if applicable.|
|ipv6_pool|text|The ID of the IPv6 address pool from which the IPv6 CIDR block is allocated.|
|network_border_group|text|The name of the unique set of Availability Zones, Local Zones, or Wavelength Zones from which AWS advertises IP addresses, for example, us-east-1-wl1-bos-wlz-1.|
