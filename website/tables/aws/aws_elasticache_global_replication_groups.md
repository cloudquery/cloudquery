# Table: aws_elasticache_global_replication_groups

This table shows data for Elasticache Global Replication Groups.

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_GlobalReplicationGroup.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|at_rest_encryption_enabled|`bool`|
|auth_token_enabled|`bool`|
|cache_node_type|`utf8`|
|cluster_enabled|`bool`|
|engine|`utf8`|
|engine_version|`utf8`|
|global_node_groups|`json`|
|global_replication_group_description|`utf8`|
|global_replication_group_id|`utf8`|
|members|`json`|
|status|`utf8`|
|transit_encryption_enabled|`bool`|