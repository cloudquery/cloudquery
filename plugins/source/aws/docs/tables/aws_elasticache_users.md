# Table: aws_elasticache_users

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_User.html

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
|access_string|String|
|authentication|JSON|
|engine|String|
|minimum_engine_version|String|
|status|String|
|user_group_ids|StringArray|
|user_id|String|
|user_name|String|