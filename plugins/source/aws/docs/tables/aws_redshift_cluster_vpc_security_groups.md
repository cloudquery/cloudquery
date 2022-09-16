
# Table: aws_redshift_cluster_vpc_security_groups
Describes the members of a VPC security group.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of aws_redshift_clusters table (FK)|
|status|text|The status of the VPC security group.|
|vpc_security_group_id|text|The identifier of the VPC security group.|
