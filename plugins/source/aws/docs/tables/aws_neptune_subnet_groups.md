# Table: aws_neptune_subnet_groups

https://docs.aws.amazon.com/neptune/latest/userguide/api-subnets.html#DescribeDBSubnetGroups

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|db_subnet_group_description|String|
|db_subnet_group_name|String|
|subnet_group_status|String|
|subnets|JSON|
|vpc_id|String|