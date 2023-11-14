# Table: aws_rds_subnet_groups

This table shows data for Amazon Relational Database Service (RDS) Subnet Groups.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBSubnetGroup.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|db_subnet_group_arn|`utf8`|
|db_subnet_group_description|`utf8`|
|db_subnet_group_name|`utf8`|
|subnet_group_status|`utf8`|
|subnets|`json`|
|supported_network_types|`list<item: utf8, nullable>`|
|vpc_id|`utf8`|