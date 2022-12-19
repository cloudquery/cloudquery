# Table: vercel_project_envs

The composite primary key for this table is (**project_id**, **id**).

## Relations

This table depends on [vercel_projects](vercel_projects.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|id (PK)|String|
|target|String|
|type|String|
|key|String|
|value|String|
|configuration_id|String|
|git_branch|String|
|edge_config_id|String|
|edge_config_token_id|String|
|decrypted|Bool|
|system|Bool|
|created_at|Timestamp|
|created_by|String|
|updated_at|Timestamp|
|updated_by|String|