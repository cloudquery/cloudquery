# Table: aws_identitystore_users



The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|identity_store_id|String|
|user_id|String|
|addresses|JSON|
|display_name|String|
|emails|JSON|
|external_ids|JSON|
|locale|String|
|name|JSON|
|nick_name|String|
|phone_numbers|JSON|
|preferred_language|String|
|profile_url|String|
|timezone|String|
|title|String|
|user_name|String|
|user_type|String|