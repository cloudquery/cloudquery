
# Table: aws_rds_subnet_group_subnets
This data type is used as a response element for the DescribeDBSubnetGroups operation. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subnet_group_cq_id|uuid|Unique CloudQuery ID of aws_rds_subnet_groups table (FK)|
|subnet_availability_zone_name|text|The name of the Availability Zone.|
|subnet_identifier|text|The identifier of the subnet.|
|subnet_outpost_arn|text|The Amazon Resource Name (ARN) of the Outpost.|
|subnet_status|text|The status of the subnet.|
