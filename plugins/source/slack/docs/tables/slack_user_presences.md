# Table: slack_user_presences



The primary key for this table is **user_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|user_id (PK)|String|
|presence|String|
|online|Bool|
|auto_away|Bool|
|manual_away|Bool|
|connection_count|Int|
|last_activity|Int|