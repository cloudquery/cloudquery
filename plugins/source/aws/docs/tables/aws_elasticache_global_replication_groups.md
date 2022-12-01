# Table: aws_elasticache_global_replication_groups

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_GlobalReplicationGroup.html

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
|at_rest_encryption_enabled|Bool|
|auth_token_enabled|Bool|
|cache_node_type|String|
|cluster_enabled|Bool|
|engine|String|
|engine_version|String|
|global_node_groups|JSON|
|global_replication_group_description|String|
|global_replication_group_id|String|
|members|JSON|
|status|String|
|transit_encryption_enabled|Bool|