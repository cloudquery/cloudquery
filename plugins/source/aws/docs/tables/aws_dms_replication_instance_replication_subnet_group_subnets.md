
# Table: aws_dms_replication_instance_replication_subnet_group_subnets
In response to a request by the DescribeReplicationSubnetGroups operation, this object identifies a subnet by its given Availability Zone, subnet identifier, and status.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|replication_instance_cq_id|uuid|Unique CloudQuery ID of aws_dms_replication_instances table (FK)|
|subnet_availability_zone_name|text|The name of the Availability Zone.|
|subnet_identifier|text|The subnet identifier.|
|subnet_status|text|The status of the subnet.|
