
# Table: aws_ec2_security_group_ip_permissions_egress_ipv6_ranges
[EC2-VPC only] Describes an IPv6 range.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_ip_permissions_egress_id|uuid|Unique ID of aws_ec2_security_group_ip_permissions_egresses table (FK)|
|cidr_ipv6|text|The IPv6 CIDR range.|
|description|text|A description for the security group rule that references this IPv6 address range.|
