
# Table: aws_elasticache_subnet_group_subnets
Represents the subnet associated with a cluster
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subnet_group_cq_id|uuid|Unique CloudQuery ID of aws_elasticache_subnet_groups table (FK)|
|subnet_availability_zone_name|text|The name of the Availability Zone.|
|subnet_identifier|text|The unique identifier for the subnet.|
|subnet_outpost_arn|text|The outpost ARN of the subnet.|
