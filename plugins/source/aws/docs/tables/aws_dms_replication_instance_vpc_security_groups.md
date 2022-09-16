
# Table: aws_dms_replication_instance_vpc_security_groups
Describes the status of a security group associated with the virtual private cloud (VPC) hosting your replication and DB instances.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|replication_instance_cq_id|uuid|Unique CloudQuery ID of aws_dms_replication_instances table (FK)|
|status|text|The status of the VPC security group.|
|vpc_security_group_id|text|The VPC security group ID.|
