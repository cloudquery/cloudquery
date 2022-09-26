# Table: aws_elasticache_users


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|