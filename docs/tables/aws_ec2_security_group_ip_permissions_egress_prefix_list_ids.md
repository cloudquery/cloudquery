
# Table: aws_ec2_security_group_ip_permissions_egress_prefix_list_ids
Describes a prefix list ID.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_ip_permissions_egress_id|uuid|Unique ID of aws_ec2_security_group_ip_permissions_egresses table (FK)|
|description|text|A description for the security group rule that references this prefix list ID.|
|prefix_list_id|text|The ID of the prefix.|
