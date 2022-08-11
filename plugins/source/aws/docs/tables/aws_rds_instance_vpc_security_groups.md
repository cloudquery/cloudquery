
# Table: aws_rds_instance_vpc_security_groups
This data type is used as a response element for queries on VPC security group membership. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_rds_instances table (FK)|
|instance_id|text|The AWS Region-unique, immutable identifier for the DB instance|
|status|text|The status of the VPC security group.|
|vpc_security_group_id|text|The name of the VPC security group.|
