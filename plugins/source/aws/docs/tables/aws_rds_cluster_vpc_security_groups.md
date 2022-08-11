
# Table: aws_rds_cluster_vpc_security_groups
This data type is used as a response element for queries on VPC security group membership. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of aws_rds_clusters table (FK)|
|status|text|The status of the VPC security group.|
|vpc_security_group_id|text|The name of the VPC security group.|
