# Table: aws_elasticache_parameter_groups

This table shows data for Elasticache Parameter Groups.

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheParameterGroup.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|cache_parameter_group_family|`utf8`|
|cache_parameter_group_name|`utf8`|
|description|`utf8`|
|is_global|`bool`|