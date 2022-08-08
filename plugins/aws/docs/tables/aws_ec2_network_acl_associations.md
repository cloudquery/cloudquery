
# Table: aws_ec2_network_acl_associations
Describes an association between a network ACL and a subnet.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|network_acl_cq_id|uuid|Unique CloudQuery ID of aws_ec2_network_acls table (FK)|
|network_acl_association_id|text|The ID of the association between a network ACL and a subnet.|
|subnet_id|text|The ID of the subnet.|
