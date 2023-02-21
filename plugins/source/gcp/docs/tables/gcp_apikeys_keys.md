# Table: gcp_apikeys_keys

https://cloud.google.com/api-keys/docs/reference/rest/v2/projects.locations.keys#Key

The composite primary key for this table is (**project_id**, **uid**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name|String|
|uid (PK)|String|
|display_name|String|
|key_string|String|
|create_time|Timestamp|
|update_time|Timestamp|
|delete_time|Timestamp|
|annotations|JSON|
|restrictions|JSON|
|etag|String|