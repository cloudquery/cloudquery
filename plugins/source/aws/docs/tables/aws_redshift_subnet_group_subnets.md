
# Table: aws_redshift_subnet_group_subnets
Describes a subnet.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subnet_group_cq_id|uuid|Unique CloudQuery ID of aws_redshift_subnet_groups table (FK)|
|subnet_availability_zone_name|text|The name of the availability zone.|
|subnet_availability_zone_supported_platforms|text[]|A list of supported platforms for orderable clusters.|
|subnet_identifier|text|The identifier of the subnet.|
|subnet_status|text|The status of the subnet.|
