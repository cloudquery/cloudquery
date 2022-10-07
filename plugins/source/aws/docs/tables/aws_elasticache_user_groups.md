# Table: aws_elasticache_user_groups



The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|arn (PK)|String|
|engine|String|
|minimum_engine_version|String|
|pending_changes|JSON|
|replication_groups|StringArray|
|status|String|
|user_group_id|String|
|user_ids|StringArray|