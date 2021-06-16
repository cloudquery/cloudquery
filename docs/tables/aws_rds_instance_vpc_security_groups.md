
# Table: aws_rds_instance_vpc_security_groups
This data type is used as a response element for queries on VPC security group membership. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid|Unique ID of aws_rds_instances table (FK)|
|status|text|The status of the VPC security group.|
|vpc_security_group_id|text|The name of the VPC security group.|
