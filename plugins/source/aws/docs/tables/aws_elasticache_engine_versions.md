# Table: aws_elasticache_engine_versions

Provides all of the details about a particular cache engine version.

The composite primary key for this table is (**account_id**, **region**, **engine**, **engine_version**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|cache_engine_description|String|
|cache_engine_version_description|String|
|cache_parameter_group_family|String|
|engine (PK)|String|
|engine_version (PK)|String|