
# Table: aws_ec2_subnets
Describes a subnet.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|assign_ipv6_address_on_creation|boolean|Indicates whether a network interface created in this subnet (including a network interface created by RunInstances) receives an IPv6 address.|
|availability_zone|text|The Availability Zone of the subnet.|
|availability_zone_id|text|The AZ ID of the subnet.|
|available_ip_address_count|integer|The number of unused private IPv4 addresses in the subnet|
|cidr_block|text|The IPv4 CIDR block assigned to the subnet.|
|customer_owned_ipv4_pool|text|The customer-owned IPv4 address pool associated with the subnet.|
|default_for_az|boolean|Indicates whether this is the default subnet for the Availability Zone.|
|map_customer_owned_ip_on_launch|boolean|Indicates whether a network interface created in this subnet (including a network interface created by RunInstances) receives a customer-owned IPv4 address.|
|map_public_ip_on_launch|boolean|Indicates whether instances launched in this subnet receive a public IPv4 address.|
|outpost_arn|text|The Amazon Resource Name (ARN) of the Outpost.|
|owner_id|text|The ID of the Amazon Web Services account that owns the subnet.|
|state|text|The current state of the subnet.|
|arn|text|The Amazon Resource Name (ARN) of the subnet.|
|id|text|The ID of the subnet.|
|tags|jsonb|Any tags assigned to the subnet.|
|vpc_id|text|The ID of the VPC the subnet is in.|
