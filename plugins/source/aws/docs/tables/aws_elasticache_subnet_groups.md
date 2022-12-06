# Table: aws_elasticache_subnet_groups

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheSubnetGroup.html

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
|cache_subnet_group_description|String|
|cache_subnet_group_name|String|
|subnets|JSON|
|supported_network_types|StringArray|
|vpc_id|String|