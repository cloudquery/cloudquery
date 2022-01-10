
# Table: aws_ec2_security_group_ip_permission_ip_ranges
Details of a cidr range associated with a security group rule
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_ip_permission_cq_id|uuid|Unique CloudQuery ID of aws_ec2_security_group_ip_permissions table (FK)|
|cidr|text|The CIDR range.|
|description|text|A description for the security group rule that references this address range.|
|cidr_type|text|IP Type: ipv4, or ipv6|
