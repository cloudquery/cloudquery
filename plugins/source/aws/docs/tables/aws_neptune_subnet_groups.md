# Table: aws_neptune_subnet_groups

This table shows data for Neptune Subnet Groups.

https://docs.aws.amazon.com/neptune/latest/userguide/api-subnets.html#DescribeDBSubnetGroups

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|db_subnet_group_description|`utf8`|
|db_subnet_group_name|`utf8`|
|subnet_group_status|`utf8`|
|subnets|`json`|
|vpc_id|`utf8`|
|db_subnet_group_arn|`utf8`|