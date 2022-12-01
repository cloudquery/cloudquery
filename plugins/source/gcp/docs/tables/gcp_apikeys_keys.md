# Table: gcp_apikeys_keys



The composite primary key for this table is (**project_id**, **uid**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|uid (PK)|String|
|name|String|
|display_name|String|
|key_string|String|
|create_time|Timestamp|
|update_time|Timestamp|
|delete_time|Timestamp|
|annotations|JSON|
|restrictions|JSON|
|etag|String|