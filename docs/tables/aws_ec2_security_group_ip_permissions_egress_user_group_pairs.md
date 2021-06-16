
# Table: aws_ec2_security_group_ip_permissions_egress_user_group_pairs
Describes a security group and AWS account ID pair.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_ip_permissions_egress_id|uuid|Unique ID of aws_ec2_security_group_ip_permissions_egresses table (FK)|
|description|text|A description for the security group rule that references this user ID group pair.|
|group_id|text|The ID of the security group.|
|group_name|text|The name of the security group.|
|peering_status|text|The status of a VPC peering connection, if applicable.|
|user_id|text|The ID of an AWS account.|
|vpc_id|text|The ID of the VPC for the referenced security group, if applicable.|
|vpc_peering_connection_id|text|The ID of the VPC peering connection, if applicable.|
