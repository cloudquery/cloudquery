
# Table: aws_ec2_subnet_ipv6_cidr_block_association_sets
Describes an IPv6 CIDR block associated with a subnet.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subnet_id|uuid|Unique ID of aws_ec2_subnets table (FK)|
|association_id|text|The association ID for the CIDR block.|
|ipv6_cidr_block|text|The IPv6 CIDR block.|
|ipv6_cidr_block_state|text|The state of a CIDR block.|
|ipv6_cidr_block_state_status_message|text|A message about the status of the CIDR block, if applicable.|
