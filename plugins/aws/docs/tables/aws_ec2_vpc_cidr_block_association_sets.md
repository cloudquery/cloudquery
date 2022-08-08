
# Table: aws_ec2_vpc_cidr_block_association_sets
Describes an IPv4 CIDR block associated with a VPC.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vpc_cq_id|uuid|Unique CloudQuery ID of aws_ec2_vpcs table (FK)|
|association_id|text|The association ID for the IPv4 CIDR block.|
|cidr_block|text|The IPv4 CIDR block.|
|cidr_block_state|text|The state of the CIDR block.|
|cidr_block_state_status_message|text|A message about the status of the CIDR block, if applicable.|
