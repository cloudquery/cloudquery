# Table: aws_elasticache_engine_versions

This table shows data for Elasticache Engine Versions.

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheEngineVersion.html

The composite primary key for this table is (**account_id**, **region**, **engine**, **engine_version**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|engine (PK)|`utf8`|
|engine_version (PK)|`utf8`|
|cache_engine_description|`utf8`|
|cache_engine_version_description|`utf8`|
|cache_parameter_group_family|`utf8`|