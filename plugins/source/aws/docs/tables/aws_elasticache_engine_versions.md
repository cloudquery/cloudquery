# Table: aws_elasticache_engine_versions

This table shows data for Elasticache Engine Versions.

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheEngineVersion.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **engine**, **engine_version**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|engine|`utf8`|
|engine_version|`utf8`|
|cache_engine_description|`utf8`|
|cache_engine_version_description|`utf8`|
|cache_parameter_group_family|`utf8`|