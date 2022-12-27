# Table: fastly_auth_tokens

https://developer.fastly.com/reference/api/auth-tokens/

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|access_token|String|
|created_at|Timestamp|
|expires_at|Timestamp|
|id (PK)|String|
|ip|String|
|last_used_at|Timestamp|
|name|String|
|scope|String|
|services|StringArray|
|user_id|String|