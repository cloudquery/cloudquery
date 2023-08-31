# Table: aws_elasticache_subnet_groups

This table shows data for Elasticache Subnet Groups.

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheSubnetGroup.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|cache_subnet_group_description|`utf8`|
|cache_subnet_group_name|`utf8`|
|subnets|`json`|
|supported_network_types|`list<item: utf8, nullable>`|
|vpc_id|`utf8`|